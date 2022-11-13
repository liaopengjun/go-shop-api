package request

type ShopCartParam struct {
	GoodsID  int64 `json:"goodsId" binding:"required"`
	GoodsNum int   `json:"goodsCount" binding:"required"`
}
