package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type ShopGoodsRouter struct {
}

// InitShopGoogsRouter 商城商品模块路由
func (s *ShopGoodsRouter) InitShopGoodsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopGoodsApi = v1.ApiGroupApp.ShopApiGroup.ShopGoodsApi
	{
		Router.GET("search", ShopGoodsApi.SearchGoodsList)     //搜索商品
		Router.POST("categories", ShopGoodsApi.GoodCategories) //商品分类
		Router.POST("goods-detail", ShopGoodsApi.GoodDetail)   //商品详情
		Router.POST("add-cart", ShopGoodsApi.AddShopCart)      //加入购物车
		Router.POST("cart-list", ShopGoodsApi.ShopCartList)    //购物车列表
	}
	return Router
}
