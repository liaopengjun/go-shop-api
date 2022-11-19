package shop

import (
	"github.com/gin-gonic/gin"
	"go-admin/config"
	"go-admin/global"
	requestcommon "go-admin/model/common/request"
	"go-admin/model/common/response"
	"go-admin/model/shop/request"
	"go.uber.org/zap"
)

type ShopGoodsApi struct {
}

// SearchGoodsList 搜索
func (goods *ShopGoodsApi) SearchGoodsList(c *gin.Context) {
	var p = new(request.GoodsParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("goodsparam fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	list, total, err := goodsService.GetSearchGoodsList(p)
	if err != nil {
		global.GA_SHOPLOG.Error("get goodslist fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, response.PageResult{
		List:  list,
		Total: total,
		Page:  p.PageNumber,
		Limit: 10,
	})
}

// GoodDetail 商品详情
func (goods *ShopGoodsApi) GoodDetail(c *gin.Context) {
	var p = new(requestcommon.GetById)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("goodsDetail fail:", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	res, err := goodsService.GetGoodsDetail(int64(p.ID))
	if err != nil {
		global.GA_SHOPLOG.Error("get goodsDetail fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, res)
}
