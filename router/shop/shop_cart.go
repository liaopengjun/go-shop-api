package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type ShopCartRouter struct {
}

func (c *ShopCartRouter) InitShopCartRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopCartApi = v1.ApiGroupApp.ShopApiGroup.ShopCartApi
	{
		Router.POST("cart-list", ShopCartApi.GetCartList)    //购物车列表
		Router.POST("cart-add", ShopCartApi.AddShopCart)     //添加购物车
		Router.POST("cart-edit", ShopCartApi.UpdateShopCart) //编辑购物车
		Router.POST("cart-amount", ShopCartApi.GetCartAmout) //编辑购物车
		Router.POST("cart-settle", ShopCartApi.Settle)       //购物车入单明细
	}
	return Router
}
