package shop

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-admin/config"
	"go-admin/global"
	requestcommon "go-admin/model/common/request"
	"go-admin/model/common/response"
	"go-admin/model/shop/request"
	GoodsRep "go-admin/model/shop/response"
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
	if len(list) <= 0 {
		list = []GoodsRep.GoodsList{}
	}
	if err != nil {
		global.GA_SHOPLOG.Error("get goodslist fail:", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	totalPage := total / 10
	response.ResponseSuccess(c, response.ShopPageResult{
		List:       list,
		TotalCount: total,
		TotalPage:  int(totalPage),
		CurrPage:   p.PageNumber,
		PageSize:   10,
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
		if errors.Is(err, GoodsRep.ErrLowerShelfError) {
			response.ResponseError(c, config.CodeLowerShelfError)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	response.ResponseSuccess(c, res)
}
