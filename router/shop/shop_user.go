package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type ShopUserRouter struct {
}

// InitShopUserRouter 商城用户模板路由
func (a *ShopUserRouter) InitShopUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopUserApi = v1.ApiGroupApp.ShopApiGroup.ShopUserApi
	{
		Router.POST("login", ShopUserApi.Login)            //登录用户
		Router.POST("out-login", ShopUserApi.OutLogin)     //退出登录
		Router.POST("register", ShopUserApi.Register)      //注册用户
		Router.POST("edit-user", ShopUserApi.EditUserInfo) //编辑用户
		Router.POST("user-info", ShopUserApi.UserInfo)     //用户详情
	}
	return Router

}
