package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-shop-api/config"
	"go-shop-api/global"
	commonRequest "go-shop-api/model/common/request"
	"go-shop-api/model/common/response"
	Request "go-shop-api/model/system/request"
	"go.uber.org/zap"
)

type SysApi struct {
}

// GetSysApiList api列表
func (a *SysApi) GetSysApiList(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(Request.ParamSysApiList)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("api列表参数有误", zap.Error(err))
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
	list, total, err := sysApiService.GetSysApiList(p)
	if err != nil {
		global.GA_LOG.Error("获取api列表失败", zap.Error(err))
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

// AddSysApi 添加api
func (a *SysApi) AddSysApi(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(Request.ParamSysApiData)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("创建api参数有误", zap.Error(err))
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
	//3.业务处理
	if err := sysApiService.AddSysApi(p); err != nil {
		global.GA_LOG.Error("创建api失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//4.返回响应
	response.ResponseSuccess(c, "创建成功")

}

// UpdateSysApi更新api
func (a *SysApi) UpdateSysApi(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(Request.ParamSysApiData)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("更新api参数有误", zap.Error(err))
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
	//3.业务处理
	if err := sysApiService.UpdateSysApi(p); err != nil {
		global.GA_LOG.Error("更新api失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//4.返回响应
	response.ResponseSuccess(c, "更新成功")
}

// DelSysApi 删除api
func (a *SysApi) DelSysApi(c *gin.Context) {
	var p = new(commonRequest.GetByIds)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("删除api参数有误", zap.Error(err))
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
	if err := sysApiService.DelSysApi(p.ID); err != nil {
		global.GA_LOG.Error("删除api有误", zap.Error(err))
		return
	}
	response.ResponseSuccess(c, "删除成功")
}

// GetSysApiAll api权限列表
func (a *SysApi) GetSysApiAll(c *gin.Context) {
	list, err := sysApiService.GetSysApiAll()
	if err != nil {
		global.GA_LOG.Error("获取api权限失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, list)
}
