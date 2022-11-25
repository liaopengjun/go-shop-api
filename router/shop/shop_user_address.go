package shop

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type ShopUserAddressRouter struct {
}

func (a *ShopUserAddressRouter) InitShopUserAddressRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var ShopUserAddressApi = v1.ApiGroupApp.ShopApiGroup.ShopUserAddressApi
	{
		Router.POST("add-address", ShopUserAddressApi.AddUserAddress)            //添加收货地址
		Router.POST("del-address", ShopUserAddressApi.DelUserAddress)            //删除收货地址
		Router.POST("edit-address", ShopUserAddressApi.EditUserAddress)          //编辑收货地址
		Router.POST("address-detail", ShopUserAddressApi.GetUserAddressInfo)     //收货地址详情
		Router.POST("address-list", ShopUserAddressApi.GetUserAddressList)       //收货地址列表
		Router.POST("default-address", ShopUserAddressApi.GetDefaultAddressInfo) //用户默认地址详情
	}
	return Router

}
