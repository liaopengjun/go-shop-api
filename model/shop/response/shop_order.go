package response

import (
	"errors"
	"time"
)

var (
	ErrCartItem         = errors.New("下单商品异常")
	ErrGoodsNonExistent = errors.New("商品不存在")
	ErrOrderLowerShelf  = errors.New("订单商品已下架")
	ErrGoodsInventory   = errors.New("商品库存不充足")
	ErrGoodsTotalPrice  = errors.New("商品价格有误")
	ErrCreateOrder      = errors.New("创建订单失败")
)

type OrderResponse struct {
	OrderId           int                 `json:"orderId"`
	OrderNo           string              `json:"orderNo"`
	TotalPrice        int                 `json:"totalPrice"`
	PayType           int                 `json:"payType"`
	OrderStatus       int                 `json:"orderStatus"`
	OrderStatusString string              `json:"orderStatusString"`
	CreateTime        time.Time           `json:"createTime"`
	OrderItemResponse []OrderItemResponse `json:"orderItemResponse"`
}

type OrderItemResponse struct {
	GoodsId       int    `json:"goodsId"`
	GoodsName     string `json:"goodsName"`
	GoodsCount    int    `json:"goodsCount"`
	GoodsCoverImg string `json:"goodsCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
}

type OrderDetailResponse struct {
	OrderNo           string              `json:"orderNo"`
	TotalPrice        int                 `json:"totalPrice"`
	PayStatus         int                 `json:"payStatus"`
	PayType           int                 `json:"payType"`
	PayTypeString     string              `json:"payTypeString"`
	PayTime           time.Time           `json:"payTime"`
	OrderStatus       int                 `json:"orderStatus"`
	OrderStatusString string              `json:"orderStatusString"`
	CreateTime        time.Time           `json:"createTime"`
	OrderItemResponse []OrderItemResponse `json:"orderItemResponse"`
}
