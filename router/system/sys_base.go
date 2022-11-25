package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-shop-api/api/v1"
)

type BaseRouter struct {
}

//InitBaseRouter 基础路由(登录)
func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login) //登录
	}
	return baseRouter
}
