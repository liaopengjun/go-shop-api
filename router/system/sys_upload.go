package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-shop-api/api/v1"
)

type UploadRoute struct {
}

// InitUploadRoute 上传文件路由
func (r *UploadRoute) InitUploadRoute(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	var uploadApi = v1.ApiGroupApp.SystemApiGroup.UploadApi
	{
		fileRouter.POST("upload", uploadApi.Upload) //上传图片
	}
}
