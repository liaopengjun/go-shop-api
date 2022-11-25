package shop

import "go-shop-api/service"

type ApiShopGroup struct {
	ShopHomeApi
	ShopUserApi
	ShopGoodsApi
	ShopOrderApi
	ShopCategoryApi
	ShopCartApi
	ShopUserAddressApi
}

var userService = service.ServiceGroupApp.ShopServiceGroup.ShopUserService
var carouselService = service.ServiceGroupApp.ShopServiceGroup.ShopCarouselService
var goodsService = service.ServiceGroupApp.ShopServiceGroup.ShopGoodsService
var categoryService = service.ServiceGroupApp.ShopServiceGroup.ShopCategoryService
var cartService = service.ServiceGroupApp.ShopServiceGroup.ShopCartService
var userAddressService = service.ServiceGroupApp.ShopServiceGroup.ShopUserAddressService
var orderService = service.ServiceGroupApp.ShopServiceGroup.ShopOrderService
