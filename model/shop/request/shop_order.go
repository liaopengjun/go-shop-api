package request

type SaveOrderParam struct {
	CartItemIds []int `json:"cartItemIds" binding:"required"`
	AddressId   int   `json:"addressId" binding:"required"`
}

type OrderPayParam struct {
	OrderNo string `json:"orderNo" binding:"required"`
	PayType string `json:"payType" binding:"required"`
}

type OrderListParam struct {
	PageNumber int `json:"pageNumber"`
	Status     int `json:"status"`
}

type OrderDetailParam struct {
	OrderNo string `json:"orderNo" binding:"required"`
}
