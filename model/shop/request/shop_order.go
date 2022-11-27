package request

type SaveOrderParam struct {
	CartItemIds []int `json:"cartItemIds" binding:"required"`
	AddressId   int   `json:"addressId" binding:"required"`
}
