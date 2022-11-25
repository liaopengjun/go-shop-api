package v1

import (
	"go-shop-api/api/v1/shop"
	"go-shop-api/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	ShopApiGroup   shop.ApiShopGroup
}

var ApiGroupApp = new(ApiGroup)
