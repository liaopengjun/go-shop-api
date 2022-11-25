package utils

import (
	"fmt"
	"go-shop-api/global"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// PathExists 文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetImageName 获取图片名称
func GetImageName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = MD5V(fileName)

	return fileName + ext
}

// GetImagePath 图片存放路径
func GetImagePath() string {
	return global.GA_CONFIG.LocalConfig.Path
}

// CheckImageExt 文件名后缀
func CheckImageExt(fileName string) bool {
	ext := path.Ext(fileName)
	for _, allowExt := range global.GA_CONFIG.LocalConfig.ImageAllowExits {
		if strings.ToUpper(string(allowExt)) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

// CheckImageSize 文件大小
func CheckImageSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	fileSize := len(content)
	if err != nil {
		return false
	}
	return fileSize <= global.GA_CONFIG.LocalConfig.ImageMaxSize
}

// CheckPermission 文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	ok, err := PathExists(dir + "/" + src)
	if !ok {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
