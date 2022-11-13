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
		Router.POST("search", ShopGoodsApi.SearchGoodsList)  //搜索商品
		Router.POST("goods-detail", ShopGoodsApi.GoodDetail) //商品详情
	}
	return Router
}
