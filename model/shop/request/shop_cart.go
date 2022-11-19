package request

type ShopCartParam struct {
	GoodsID  int64 `json:"goodsId" binding:"required"`
	GoodsNum int   `json:"goodsCount" binding:"required"`
}

type ShopCartListParam struct {
	PageNumber int `json:"pageNumber"`
}

type ShopEditCartParam struct {
	CartItemID int `json:"cartItemId" binding:"required"`
	GoodsCount int `json:"goodsCount" binding:"required"`
}

type CartItemIDsParam struct {
	CartItemIds string `json:"cartItemIds"`
}
