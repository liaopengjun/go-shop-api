package shop

import (
	"go-admin/model/shop"
	"go-admin/model/shop/request"
	"go-admin/model/shop/response"
	"go-admin/utils"
)

type ShopGoodsService struct {
}

func (g *ShopGoodsService) GetGoodsList(param *request.GoodsParam) (data response.HomeGoodsData, total int64, err error) {
	goodsAll, total, err := shop.GetGoodsList("home", param)
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

func (g *ShopGoodsService) GetSearchGoodsList(param *request.GoodsParam) (data []response.GoodsList, total int64, err error) {
	goodsAll, total, err := shop.GetGoodsList("", param)
	for _, v := range goodsAll {
		res := response.GoodsList{
			GoodsId:       v.GoodsId,
			GoodsName:     utils.SubStrLen(v.GoodsName, 30),
			GoodsIntro:    utils.SubStrLen(v.GoodsIntro, 30),
			GoodsCoverImg: v.GoodsCoverImg,
			SellingPrice:  v.SellingPrice,
		}
		data = append(data, res)
	}
	return
}

func (g *ShopGoodsService) GetGoodsDetail(id int64) (data response.HomeGoodsDetail, err error) {
	s, err := shop.GetGoodsDetail(id)
	if s.GoodsSellStatus != 0 {
		return response.HomeGoodsDetail{}, response.ErrLowerShelfError
	}
	data = response.HomeGoodsDetail{
		GoodsId:            s.GoodsId,
		GoodsName:          utils.SubStrLen(s.GoodsName, 1000),
		GoodsIntro:         s.GoodsIntro,
		GoodsDetailContent: s.GoodsDetailContent,
		GoodsCoverImg:      s.GoodsCoverImg,
		SellingPrice:       s.SellingPrice,
		Tag:                s.Tag,
	}
	var list []string
	data.GoodsCarouselList = append(list, s.GoodsCarousel)
	return
}
