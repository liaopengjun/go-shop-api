package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-shop-api/api/v1"
)

type ShopBasicRouter struct {
}

// InitShopBasicRouter 商城基础模块路由
func (r *ShopBasicRouter) InitShopBasicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopHomeApi = v1.ApiGroupApp.ShopApiGroup.ShopHomeApi
	var ShopUserApi = v1.ApiGroupApp.ShopApiGroup.ShopUserApi
	{
		Router.POST("index", ShopHomeApi.IndexInfo)    //首页
		Router.POST("shop-info", ShopHomeApi.ShopInfo) //商城信息
		Router.POST("login", ShopUserApi.Login)        //登录用户
		Router.POST("register", ShopUserApi.Register)  //注册用户

	}
	return Router

}
