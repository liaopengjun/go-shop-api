package shop

import (
	"github.com/gin-gonic/gin"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/model/shop"
	"go-admin/model/shop/request"
	"go.uber.org/zap"
	"time"
)

type ShopCartApi struct {
}

// GetCartList
func (cart *ShopCartApi) GetCartList(c *gin.Context) {
	var p = new(request.ShopCartListParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("get cart-list param fail :", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	shop_userid, _ := c.Get("shop_userid")
	list, _, err := cartService.GetCartList(shop_userid.(uint), p.PageNumber)
	if err != nil {
		global.GA_SHOPLOG.Error("add cart fail :", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, list)
}

// AddShopCart 添加购物车
func (cart *ShopCartApi) AddShopCart(c *gin.Context) {
	var p = new(request.ShopCartParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("add cart param fail :", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	shop_userid, _ := c.Get("shop_userid")
	err := cartService.AddCart(shop_userid.(uint), p.GoodsID, p.GoodsNum)
	if err != nil {
		global.GA_SHOPLOG.Error("add cart fail :", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "添加购物车成功")
}

// UpdateShopCart 购物车
func (cart *ShopCartApi) UpdateShopCart(c *gin.Context) {
	var p = new(request.ShopEditCartParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("edit cart param fail :", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	cartData := shop.ShopCartItem{
		CartItemId: p.CartItemID,
		GoodsCount: p.GoodsCount,
		UpdateTime: time.Now(),
	}
	err := shop.UpdateCart(cartData)
	if err != nil {
		global.GA_SHOPLOG.Error("edit cart fail :", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "修改成功")
}

func (cart *ShopCartApi) GetCartAmout(c *gin.Context) {
	shop_userid, _ := c.Get("shop_userid")
	total, err := cartService.GetCartAmout(shop_userid.(uint))
	if err != nil {
		global.GA_SHOPLOG.Error("get userCartCount fail :", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, total)
}

//Settle 购物车入单明细
func (cart *ShopCartApi) Settle(c *gin.Context) {

}