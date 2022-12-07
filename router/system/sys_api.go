package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-shop-api/api/v1"
)

type ApiRouter struct {
}

// InitApiRouter api模块路由
func (a *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	sysApiRouter := Router.Group("sys_api")
	var SysApi = v1.ApiGroupApp.SystemApiGroup.SysApi
	{
		sysApiRouter.POST("addSysApi", SysApi.AddSysApi)         // 创建api
		sysApiRouter.POST("delSysApi", SysApi.DelSysApi)         // 删除api
		sysApiRouter.POST("getSysApiList", SysApi.GetSysApiList) // api列表
		sysApiRouter.POST("updateSysApi", SysApi.UpdateSysApi)   // 更新api
		sysApiRouter.POST("getSysApiAll", SysApi.GetSysApiAll)   // api权限
	}
}
