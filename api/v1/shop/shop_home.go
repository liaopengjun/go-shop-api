package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/model/common/response"
	"go-admin/model/shop/request"
)

type ShopHomeApi struct {
}

// IndexInfo 首页
func (b *ShopHomeApi) IndexInfo(c *gin.Context) {
	// 轮播图
	bannerList, err := carouselService.GetCarouselList()
	if err != nil {
		return
	}
	var param = new(request.GoodsParam)
	goodsList, err := goodsService.GetGoodsList(param)
	if err != nil {
		return
	}
	indexResult := make(map[string]interface{})
	indexResult["carousels"] = bannerList
	indexResult["hotGoodses"] = goodsList.HomeHotGoods
	indexResult["newGoodses"] = goodsList.HomeNewGoods
	indexResult["recommendGoodses"] = goodsList.HomeRecommendGoods
	response.ResponseSuccess(c, indexResult)
}

// ShopInfo 商信息
func (b *ShopHomeApi) ShopInfo(g *gin.Context) {
	fmt.Println("商城信息")
}
