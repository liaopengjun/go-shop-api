package shop

import (
	"github.com/gin-gonic/gin"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/common/response"
	"go.uber.org/zap"
)

type ShopCategoryApi struct {
}

// GetCategoryList 搜索
func (category *ShopCategoryApi) GetCategoryList(c *gin.Context) {
	categoryList, err := categoryService.GetGoodsCategoryList()
	if err != nil {
		global.GA_SHOPLOG.Error("get category fail", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, categoryList)
}
