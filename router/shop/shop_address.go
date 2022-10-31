package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type ShopAddressRouter struct {
}

// InitShopAddressRouter 商城地址模块路由
func (r *ShopBasicRouter) InitShopAddressRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopBasicApi = v1.ApiGroupApp.ShopApiGroup.ShopAddressApi
	{
		Router.POST("add-address", ShopBasicApi.AddAddress)      //添加地址
		Router.POST("del-address", ShopBasicApi.DelAddress)      //删除地址
		Router.POST("edit-address", ShopBasicApi.EditAddress)    //更新地址
		Router.POST("address-list", ShopBasicApi.GetAddressList) //地址列表
		Router.POST("address-info", ShopBasicApi.GetAddressInfo) //地址详情
	}
	return Router

}
