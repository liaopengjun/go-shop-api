package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/common/response"
	Request "go-admin/model/system/request"
	Response "go-admin/model/system/response"
	"go.uber.org/zap"
)

type AuthorityApi struct {
}

// CreateAuthority 创建角色
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(Request.ParamAuthorityData)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("创建角色参数有误", zap.Error(err))
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
	if err := authService.AddAuthority(p); err != nil {
		global.GA_LOG.Error("创建角色失败", zap.Error(err))
		if errors.Is(err, Response.ErrorAuthExit) {
			response.ResponseError(c, config.CodeAuthExist)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	//4.返回响应
	response.ResponseSuccess(c, "创建成功")

}

// CreateAuthority 删除角色
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var p = new(Request.GetAuthorityId)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("删除角色参数有误", zap.Error(err))
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
	if err := authService.DelAuthority(p.AuthorityId); err != nil {
		global.GA_LOG.Error("删除角色有误", zap.Error(err))
		if errors.Is(err, Response.ErrorAuthChildExit) {
			response.ResponseError(c, config.CodeMenuChildExist)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	response.ResponseSuccess(c, "删除成功")
}

// UpdateAuthority 更新角色
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(Request.ParamAuthorityData)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("更新角色参数有误", zap.Error(err))
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
	if err := authService.UpdateAuthority(p); err != nil {
		global.GA_LOG.Error("更新角色失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//4.返回响应
	response.ResponseSuccess(c, "更新成功")
}

// GetAuthorityList 角色列表
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	//1.定义列表请求结构体
	var p = new(Request.ParamAuthorityList)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("角色列表参数有误", zap.Error(err))
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
	list, total, err := authService.GetAuthorityList(p)
	if err != nil {
		global.GA_LOG.Error("获取菜单列表失败", zap.Error(err))
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

// SetAuthority 角色权限
func (a *AuthorityApi) SetAuthority(c *gin.Context) {
	var p = new(Request.AddMenuAuthorityInfo)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("设置角色菜单参数有误", zap.Error(err))
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
	if err := authService.SetAuthority(p.Menus, p.AuthorityId); err != nil {
		global.GA_LOG.Error("设置角色权限失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "设置成功")
}

// GetAuthorityInfo 获取权限详情
func (a *AuthorityApi) GetAuthorityInfo(c *gin.Context) {
	var p = new(Request.GetAuthorityId)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("查询角色参数有误", zap.Error(err))
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
	list := authService.GetAuthorityInfo(p.AuthorityId)
	response.ResponseSuccess(c, response.PageResult{List: list})

}

// UpdateAuthoritStatus 更新角色状态
func (a *AuthorityApi) UpdateAuthorityStatus(c *gin.Context) {
	var p = new(Request.GetAuthorityId)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("更新状态参数有误", zap.Error(err))
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
	if err := authService.UpdateAuthorityStatus(p); err != nil {
		global.GA_LOG.Error("更新状态错误", zap.Error(err))
		if errors.Is(err, Response.ErrorAuthApiExit) {
			response.ResponseError(c, config.CodeAuthApiExit)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
	}
	response.ResponseSuccess(c, "设置成功")
}

// SetAuthApi 设置角色api
func (a *AuthorityApi) SetAuthApi(c *gin.Context) {
	var p = new(Request.CasbinInReceive)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("设置角色api参数有误", zap.Error(err))
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
	//业务逻辑
	if err := casbinService.UpdateCasbin(p.AuthorityId, p.CasbinInfos); err != nil {
		global.GA_LOG.Error("更新api错误", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//返回状态
	response.ResponseSuccess(c, "设置成功")
}

// GetAuthApi 获取角色API列表
func (a *AuthorityApi) GetAuthApi(c *gin.Context) {
	var p = new(Request.GetAuthorityId)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("角色api列表参数有误", zap.Error(err))
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
	//业务查询
	list := casbinService.GetPolicyPathByAuthorityId(p.AuthorityId)
	//返回数据
	response.ResponseSuccess(c, list)
}
