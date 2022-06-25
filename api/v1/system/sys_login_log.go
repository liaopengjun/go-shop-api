package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-admin/config"
	"go-admin/global"
	commonRequest "go-admin/model/common/request"
	"go-admin/model/common/response"
	"go-admin/model/system/request"
	"go.uber.org/zap"
)

type LoginLogApi struct {
}

// LoginLogList日志列表
func (l *LoginLogApi) GetLoginLogList(c *gin.Context) {
	var p = new(request.ParamLoginLogList)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("查询登陆log参数有误", zap.Error(err))
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	//3.处理业务
	list, total, err := loginLogService.GetLoginLogList(p)
	if err != nil {
		global.GA_LOG.Error("获取日志列表失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//4.处理返回
	response.ResponseSuccess(c, response.PageResult{
		List:  list,
		Total: total,
		Page:  int(p.Page),
		Limit: int(p.Limit),
	})
}

// DelLoginLog 删除登陆日志
func (l *LoginLogApi) DelLoginLog(c *gin.Context) {
	var p = new(commonRequest.GetByIds)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("删除log参数有误", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	err := loginLogService.DelLoginLog(p.ID)
	if err != nil {
		global.GA_LOG.Error("删除登陆日志失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "删除成功")
}
