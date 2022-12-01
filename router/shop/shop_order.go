package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-shop-api/api/v1"
)

type ShopOrderRouter struct {
}

// InitShopOrderRouter 商城订单路由
func (s *ShopOrderRouter) InitShopOrderRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopOrderApi = v1.ApiGroupApp.ShopApiGroup.ShopOrderApi
	{
		Router.POST("create-order", ShopOrderApi.CreateOrder)      //创建订单
		Router.POST("cancel-order", ShopOrderApi.CancelOrder)      //取消订单
		Router.POST("order-pay", ShopOrderApi.OrderPay)            //订单支付
		Router.POST("order-list", ShopOrderApi.OrderList)          //订单列表
		Router.POST("order-detail", ShopOrderApi.GetOrderDetail)   //订单详情
		Router.POST("confirm-order", ShopOrderApi.ConfirmTheGoods) //确认收货
	}
	return Router

}
