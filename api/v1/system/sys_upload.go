package system

import (
	"github.com/gin-gonic/gin"
	"go-shop-api/config"
	"go-shop-api/global"
	"go-shop-api/model/common/response"
	"go.uber.org/zap"
)

type UploadApi struct {
}

// Upload 上传图片
func (u *UploadApi) Upload(c *gin.Context) {
	_, image, err := c.Request.FormFile("file")
	if err != nil {
		global.GA_LOG.Error("角色api列表参数有误", zap.Error(err))
		response.ResponseError(c, config.CodeFileError)
		return
	}
	err, filePath := uploadService.UploadFile(image)
	if err != nil {
		global.GA_LOG.Error("上传文件失败", zap.Error(err))
	}
	//4.返回响应
	response.ResponseSuccess(c, filePath)
}
