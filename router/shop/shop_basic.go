package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type ShopBasicRouter struct {
}

// InitShopBasicRouter 商城基础模块路由
func (r *ShopBasicRouter) InitShopBasicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopBasicApi = v1.ApiGroupApp.ShopApiGroup.ShopBasicApi
	{
		Router.POST("index", ShopBasicApi.IndexInfo) //首页
	}
	return Router

}
