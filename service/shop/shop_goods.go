package shop

import (
	"go-admin/model/shop"
	"go-admin/model/shop/request"
	"go-admin/model/shop/response"
	"go-admin/utils"
)

type ShopGoodsService struct {
}

func (g *ShopGoodsService) GetGoodsList(param *request.GoodsParam) (data response.HomeGoodsData, err error) {
	goodsAll, err := shop.GetGoodsList("home", param)
	//分类处理结果
	var goodsHotList []response.HomeGoodsDetail
	var goodsNewList []response.HomeGoodsDetail
	var goodsRecommendList []response.HomeGoodsDetail
	for _, v := range goodsAll {
		res := response.HomeGoodsDetail{
			GoodsId:       v.GoodsId,
			GoodsName:     utils.SubStrLen(v.GoodsName, 30),
			GoodsIntro:    utils.SubStrLen(v.GoodsIntro, 30),
			GoodsCoverImg: v.GoodsCoverImg,
			SellingPrice:  v.SellingPrice,
			Tag:           v.Tag,
		}
		switch v.GoodsType {
		case 1:
			goodsNewList = append(goodsNewList, res)
		case 2:
			goodsHotList = append(goodsHotList, res)
		case 3:
			goodsRecommendList = append(goodsRecommendList, res)
		default:
			continue
		}
	}
	data.HomeHotGoods = goodsHotList
	data.HomeNewGoods = goodsNewList
	data.HomeRecommendGoods = goodsRecommendList
	return
}
