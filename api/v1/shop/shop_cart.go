package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/model/shop/request"
	"go.uber.org/zap"
)

type ShopCartApi struct {
}

// GetCartList
func (cart *ShopCartApi) GetCartList(c *gin.Context) {
	fmt.Println("购物车列表")
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
	fmt.Println("购物车")
}
