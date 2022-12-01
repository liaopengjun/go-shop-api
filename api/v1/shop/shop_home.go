package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-shop-api/global"
	"go-shop-api/model/common/response"
	"go-shop-api/model/shop/request"
	"go.uber.org/zap"
)

type ShopHomeApi struct {
}

// IndexInfo 首页
func (b *ShopHomeApi) IndexInfo(c *gin.Context) {
	bannerList, err := carouselService.GetCarouselList()
	if err != nil {
		global.GA_SHOPLOG.Error(" Banner fail", zap.Error(err))
		return
	}
	var param = new(request.GoodsParam)
	goodsList, _, err := goodsService.GetGoodsList(param)
	if err != nil {
		global.GA_SHOPLOG.Error(" Home stock fail", zap.Error(err))
		return
	}
	indexResult := make(map[string]interface{})
	indexResult["carousels"] = bannerList
	indexResult["hotGoodses"] = goodsList.HomeHotGoods
	indexResult["newGoodses"] = goodsList.HomeNewGoods
	indexResult["recommendGoodses"] = goodsList.HomeRecommendGoods
	response.ResponseSuccess(c, goodsList)
}

// ShopInfo 商信息
func (b *ShopHomeApi) ShopInfo(g *gin.Context) {
	fmt.Println("商城信息")
}
