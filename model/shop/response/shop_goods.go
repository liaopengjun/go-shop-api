package response

import "errors"

var (
	ErrLowerShelfError = errors.New("商品已下架")
)

type HomeGoodsData struct {
	HomeHotGoods       []HomeGoodsDetail
	HomeNewGoods       []HomeGoodsDetail
	HomeRecommendGoods []HomeGoodsDetail
}

type HomeGoodsDetail struct {
	GoodsId            int      `json:"goodsId"`
	GoodsName          string   `json:"goodsName"`
	GoodsIntro         string   `json:"goodsIntro"`
	GoodsCoverImg      string   `json:"goodsCoverImg"`
	GoodsDetailContent string   `json:"goodsDetailContent"  `
	SellingPrice       int      `json:"sellingPrice"`
	Tag                string   `json:"tag"`
	GoodsCarouselList  []string `json:"goodsCarouselList" `
}

type GoodsList struct {
	GoodsId       int    `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	GoodsIntro    string `json:"goodsIntro"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
}
