package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-shop-api/config"
	"go-shop-api/global"
	"go-shop-api/model/common/request"
	"go-shop-api/model/common/response"
	"go-shop-api/model/system"
	Request "go-shop-api/model/system/request"
	Response "go-shop-api/model/system/response"
	"go-shop-api/utils"
	"go.uber.org/zap"
	"strings"
)

type MenuApi struct {
}

func (a *MenuApi) GetMenu(c *gin.Context) {

}
func (a *MenuApi) AddMenu(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(system.SysMenu)

	//2.校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("添加菜单参数有误", zap.Error(err))
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
	if err := menuService.AddMenu(p); err != nil {
		global.GA_LOG.Error("创建菜单失败", zap.Error(err))
		if errors.Is(err, Response.ErrorMenuExit) {
			response.ResponseError(c, config.CodeMenuExist)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	//4.返回响应
	response.ResponseSuccess(c, "创建菜单成功")
}

func (a *MenuApi) DeleteMenu(c *gin.Context) {
	//1.获取注册请求参数结构体
	GetById := new(request.GetById)
	if err := c.ShouldBindJSON(GetById); err != nil {
		global.GA_LOG.Error("删除菜单参数有误", zap.Error(err))
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
	//2.处理业务
	if err := menuService.DelMenu(int(GetById.ID)); err != nil {
		global.GA_LOG.Error("删除菜单失败", zap.Error(err))
		if errors.Is(err, Response.ErrorChildExit) {
			response.ResponseError(c, config.CodeMenuChildExist)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	//3.处理响应
	response.ResponseSuccess(c, "删除菜单成功")
}

func (a *MenuApi) UpdateMenu(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(system.SysMenu)
	//2.校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("添加菜单参数有误", zap.Error(err))
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
	if err := menuService.UpdateMenu(p); err != nil {
		global.GA_LOG.Error("更新菜单失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//4.响应返回
	response.ResponseSuccess(c, "更新菜单成功")
}

func (a *MenuApi) GetMenuList(c *gin.Context) {
	//1.定义列表请求结构体
	var p = new(Request.ParamMenuList)
	//2.参数校验
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("菜单列表参数有误", zap.Error(err))
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
	list, total, err := menuService.GetMenuList(p)
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

// getMenuTreeList 菜单树形结构
func (a *MenuApi) GetMenuTree(c *gin.Context) {
	list, total, err := menuService.GetMenuTreeList()
	if err != nil {
		global.GA_LOG.Error("获取菜单列表失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, response.PageResult{
		List:  list,
		Total: total,
	})
}

func (a *MenuApi) AddMenuAuthority(c *gin.Context) {}

// GetMenuInfo 获取菜单详情
func (a *MenuApi) GetMenuInfo(c *gin.Context) {
	//1.获取注册请求参数结构体
	GetById := new(request.GetById)
	if err := c.ShouldBindJSON(GetById); err != nil {
		global.GA_LOG.Error("菜单详情参数有误", zap.Error(err))
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
	// 2.业务处理
	list := menuService.GetMenuInfo(int(GetById.ID))
	//4.处理返回
	response.ResponseSuccess(c, response.PageResult{List: list})
}

// GetAuthorityMenuList 获取角色权限菜单
func (a *MenuApi) GetAuthorityMenuList(c *gin.Context) {
	var p = new(Request.GetAuthorityId)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("获取角色菜单参数有误", zap.Error(err))
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
	err, list := menuService.GetAuthorityMenuList(p)
	if err != nil {
		global.GA_LOG.Error("查询角色权限失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, response.PageResult{List: list})
}

// GetUserMenuList 获取用户菜单
func (a *MenuApi) GetUserMenuList(c *gin.Context) {
	//菜单权限
	authorityId := utils.GetUserAuthorityId(c)
	menuList, err := menuService.GetUserMenuList(authorityId)
	if err != nil {
		global.GA_LOG.Error("获取用户侧边栏菜单有误", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//获取用户按钮权限
	btnList := casbinService.GetPolicyPathByAuthorityId(authorityId)
	var BtnList []string
	for _, v := range btnList {
		cutBtn := v.Path[1:]
		btn := strings.Replace(cutBtn, "/", ":", -1)
		BtnList = append(BtnList, btn)
	}
	response.ResponseSuccess(c, Response.UserResult{
		MenuList: menuList,
		BtnList:  BtnList,
	})

}
