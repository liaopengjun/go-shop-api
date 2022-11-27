package response

import "errors"

var (
	ErrCartItem         = errors.New("下单商品异常")
	ErrGoodsNonExistent = errors.New("商品不存在")
	ErrOrderLowerShelf  = errors.New("订单商品已下架")
	ErrGoodsInventory   = errors.New("商品库存不充足")
	ErrGoodsTotalPrice  = errors.New("商品价格有误")
)
