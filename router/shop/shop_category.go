package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type ShopCategoryRouter struct {
}

// InitShopCategoryRouter 商城分类模块路由
func (c *ShopCategoryRouter) InitShopCategoryRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopCategoryApi = v1.ApiGroupApp.ShopApiGroup.ShopCategoryApi
	{
		Router.POST("categories", ShopCategoryApi.GetCategoryList) //商品分类
	}
	return Router
}
