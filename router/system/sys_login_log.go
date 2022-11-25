package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-shop-api/api/v1"
)

type LoginLogRouter struct {
}

// InitLoginLogRouter 登录日志路由
func (s *UserRouter) InitLoginLogRouter(Router *gin.RouterGroup) {
	logRouter := Router.Group("log")
	var LogApi = v1.ApiGroupApp.SystemApiGroup.LoginLogApi
	{
		logRouter.POST("getLoginLogList", LogApi.GetLoginLogList) // 登陆日志列表
		logRouter.POST("delLoginLog", LogApi.DelLoginLog)         // 删除登陆日志
	}
}
