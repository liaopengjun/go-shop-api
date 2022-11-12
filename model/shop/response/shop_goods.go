package response

type HomeGoodsData struct {
	HomeHotGoods       []HomeGoodsDetail
	HomeNewGoods       []HomeGoodsDetail
	HomeRecommendGoods []HomeGoodsDetail
}

type HomeGoodsDetail struct {
	GoodsId       int    `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	GoodsIntro    string `json:"goodsIntro"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
	Tag           string `json:"tag"`
}

type GoodsList struct {
	GoodsId       int    `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	GoodsIntro    string `json:"goodsIntro"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
}
