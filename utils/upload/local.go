package upload

import (
	"errors"
	"go-admin/global"
	"go-admin/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct {
}

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V(name)

	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext

	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.GA_CONFIG.LocalConfig.Path, os.ModePerm)
	if mkdirErr != nil {
		global.GA_LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.GA_CONFIG.LocalConfig.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.GA_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GA_LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GA_LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

func (*Local) DeleteFile(file_path string) error {
	//p := global.GA_CONFIG.LocalConfig.Path + "/" + key
	if strings.Contains(file_path, global.GA_CONFIG.LocalConfig.Path) {
		if err := os.Remove(file_path); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
