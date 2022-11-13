package shop

import "go-admin/service"

type ApiShopGroup struct {
	ShopHomeApi
	ShopUserApi
	ShopGoodsApi
	ShopOrderApi
	ShopAddressApi
	ShopCategoryApi
	ShopCartApi
}

var userService = service.ServiceGroupApp.ShopServiceGroup.ShopUserService
var carouselService = service.ServiceGroupApp.ShopServiceGroup.ShopCarouselService
var goodsService = service.ServiceGroupApp.ShopServiceGroup.ShopGoodsService
var categoryService = service.ServiceGroupApp.ShopServiceGroup.ShopCategoryService
var cartService = service.ServiceGroupApp.ShopServiceGroup.ShopCartService
