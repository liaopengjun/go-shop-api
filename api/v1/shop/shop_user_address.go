package shop

import (
	"github.com/gin-gonic/gin"
	"go-shop-api/config"
	"go-shop-api/global"
	requestCom "go-shop-api/model/common/request"
	"go-shop-api/model/common/response"
	"go-shop-api/model/shop/request"
	"go.uber.org/zap"
)

type ShopUserAddressApi struct {
}

// AddUserAddress 添加地址
func (s *ShopUserAddressApi) AddUserAddress(c *gin.Context) {
	var p = new(request.AddUserAddressParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("add userAddress param fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	userId, _ := c.Get("shop_userid")
	err := userAddressService.AddUserAddress(userId.(uint), p)
	if err != nil {
		global.GA_LOG.Error("add userAddress res fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "保存成功")
}

// GetUserAddressList 查询收货地址
func (s *ShopUserAddressApi) GetUserAddressList(c *gin.Context) {
	userId, _ := c.Get("shop_userid")
	list, err := userAddressService.GetUserAddressList(userId.(uint))
	if err != nil {
		global.GA_LOG.Error("add userAddress res fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, list)
}

// GetDefaultAddressInfo 用户创建订单默认地址
func (s *ShopUserAddressApi) GetDefaultAddressInfo(c *gin.Context) {
	userId, _ := c.Get("shop_userid")
	res, err := userAddressService.GetDefaultAddressInfo(userId.(uint))
	if err != nil {
		global.GA_LOG.Error("get default userAddress res fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, res)
}

// DelUserAddress 删除地址
func (s *ShopUserAddressApi) DelUserAddress(c *gin.Context) {
	var p = new(requestCom.GetById)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("del userAddress param fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	err := userAddressService.DelUserAddress(p.ID)
	if err != nil {
		global.GA_LOG.Error("del userAddress res fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "删除成功")

}

// EditUserAddress 编辑地址
func (s *ShopUserAddressApi) EditUserAddress(c *gin.Context) {
	var p = new(request.EditUserAddressParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("del userAddress param fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	userId, _ := c.Get("shop_userid")
	err := userAddressService.EditUserAddress(userId.(uint), p)
	if err != nil {
		global.GA_LOG.Error("del userAddress res fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "编辑成功")
}

// GetUserAddressInfo 获取用户地址
func (s *ShopUserAddressApi) GetUserAddressInfo(c *gin.Context) {
	var p = new(requestCom.GetById)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("del userAddress param fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	userId, _ := c.Get("shop_userid")
	res, err := userAddressService.GetUserAddressInfo(userId.(uint), p.ID)
	if err != nil {
		global.GA_LOG.Error("del userAddress res fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, res)
}
