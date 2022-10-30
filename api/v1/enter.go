package v1

import (
	"go-admin/api/v1/shop"
	"go-admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	ShopApiGroup   shop.ApiShopGroup
}

var ApiGroupApp = new(ApiGroup)
