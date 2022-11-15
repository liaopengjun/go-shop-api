package response

type CartItemResponse struct {
	CartItemId    int    `json:"cartItemId"`
	GoodsId       int64  `json:"goodsId"`
	GoodsCount    int    `json:"goodsCount"`
	GoodsName     string `json:"goodsName"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
}
