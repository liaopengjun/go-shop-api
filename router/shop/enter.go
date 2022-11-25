package shop

type RouterGroup struct {
	ShopBasicRouter
	ShopUserRouter
	ShopGoodsRouter
	ShopOrderRouter
	ShopCategoryRouter
	ShopCartRouter
	ShopUserAddressRouter
}
