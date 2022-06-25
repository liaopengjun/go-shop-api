package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type UploadRoute struct {
}

func (r UploadRoute) InitUploadRoute(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file")
	var uploadApi = v1.ApiGroupApp.SystemApiGroup.UploadApi
	{
		fileRouter.POST("upload", uploadApi.Upload) //上传图片
	}
}
