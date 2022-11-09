package shop

import "go-admin/service"

type ApiShopGroup struct {
	ShopHomeApi
	ShopUserApi
	ShopGoodsApi
	ShopOrderApi
	ShopAddressApi
}

var userService = service.ServiceGroupApp.ShopServiceGroup.UserService
var carouselService = service.ServiceGroupApp.ShopServiceGroup.CarouselService
var goodsService = service.ServiceGroupApp.ShopServiceGroup.ShopGoodsService
