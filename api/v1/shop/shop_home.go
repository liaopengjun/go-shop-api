package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/model/common/response"
)

type ShopHomeApi struct {
}

// IndexInfo 首页
func (b *ShopHomeApi) IndexInfo(c *gin.Context) {

	bannerList, err := carouselService.GetCarouselList()
	if err != nil {
		return
	}
	indexResult := make(map[string]interface{})
	indexResult["carousels"] = bannerList
	indexResult["hotGoodses"] = ""
	indexResult["newGoodses"] = ""
	indexResult["recommendGoodses"] = ""
	response.ResponseSuccess(c, indexResult)

	fmt.Println(bannerList)
	fmt.Println("首页")
}

// ShopInfo 商信息
func (b *ShopHomeApi) ShopInfo(g *gin.Context) {
	fmt.Println("商城信息")
}
