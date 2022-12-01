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
	var p = new(request.OrderPayParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("pay order fail:", zap.Error(err))
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
	if global.GA_CONFIG.UserPay {
		var responeData string
		res, err := orderService.OrderPay2(p)
		if err != nil {
			global.GA_LOG.Error("pay order fail:", zap.Error(err))
			response.ResponseError(c, config.CodePayOrderError)
			return
		}
		responeData = res["payUrl"].(string)
		response.ResponseSuccess(c, responeData)
	} else {
		err := orderService.OrderPay(p)
		if err != nil {
			global.GA_LOG.Error("pay order fail:", zap.Error(err))
			response.ResponseError(c, config.CodePayOrderError)
			return
		}
		response.ResponseSuccess(c, "支付成功")

	}
}

// OrderList 订单列表
func (o *ShopOrderApi) OrderList(c *gin.Context) {
	var p = new(request.OrderListParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("get order_list fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	if p.PageNumber <= 0 {
		p.PageNumber = 1
	}

	userId, _ := c.Get("shop_userid")
	list, total, err := orderService.GetOrderList(p, userId.(uint))
	if err != nil {
		global.GA_LOG.Error("get order-list fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	var PageSize int64 = 5
	totalPage := total / PageSize
	response.ResponseSuccess(c, response.ShopPageResult{
		List:       list,
		TotalCount: total,
		TotalPage:  int(totalPage),
		CurrPage:   p.PageNumber,
		PageSize:   int(PageSize),
	})
}

// GetOrderDetail 订单详情
func (o *ShopOrderApi) GetOrderDetail(c *gin.Context) {
	var p = new(request.OrderDetailParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("get order_list fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	res, err := orderService.GetOrderDetail(p.OrderNo)
	if err != nil {
		global.GA_LOG.Error("get order-detail fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, res)

}

// CancelOrder 取消订单
func (o *ShopOrderApi) CancelOrder(c *gin.Context) {
	var p = new(request.OrderDetailParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error(" CancelOrder param fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	err := orderService.CancelOrder(p.OrderNo)
	if err != nil {
		global.GA_LOG.Error(" CancelOrder fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "取消订单成功")

}

// ConfirmTheGoods 确认收货
func (o *ShopOrderApi) ConfirmTheGoods(c *gin.Context) {
	var p = new(request.OrderDetailParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("ConfirmTheGoods param fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	err := orderService.ConfirmTheGoods(p.OrderNo)
	if err != nil {
		global.GA_LOG.Error("ConfirmTheGoods fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "确认收货成功")
}
