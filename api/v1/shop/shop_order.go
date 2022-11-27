package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-shop-api/config"
	"go-shop-api/global"
	"go-shop-api/model/common/response"
	"go-shop-api/model/shop/request"
	"go-shop-api/utils"
	"go.uber.org/zap"
)

type ShopOrderApi struct {
}

// CreateOrder 创建订单
func (o *ShopOrderApi) CreateOrder(c *gin.Context) {
	var p = new(request.SaveOrderParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("create order fail:", zap.Error(err))
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, utils.RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	userId, _ := c.Get("shop_userid")
	orderCode, err := orderService.CreateOrder(p, userId.(uint))
	if err != nil {
		global.GA_LOG.Error("create order fail:", zap.Error(err))
		response.ResponseError(c, config.CodeCreateOrderError)
		return
	}
	response.ResponseSuccess(c, orderCode)
}

// OrderPay 订单支付
func (o *ShopOrderApi) OrderPay(c *gin.Context) {

}

// OrderList 订单列表
func (o *ShopOrderApi) OrderList(c *gin.Context) {

}

// CancelOrder 取消订单
func (o *ShopOrderApi) CancelOrder(c *gin.Context) {

}

// ConfirmTheGoods 确认收货
func (o *ShopOrderApi) ConfirmTheGoods(c *gin.Context) {

}
