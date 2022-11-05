package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/model/shop/request"
	"go-admin/utils"
	"go.uber.org/zap"
)

type ShopUserApi struct {
}

// Login 用户登录
func (u *ShopUserApi) Login(c *gin.Context) {

}

// OutLogin 退出登录
func (u *ShopUserApi) OutLogin(c *gin.Context) {

}

// Register 注册
func (u *ShopUserApi) Register(c *gin.Context) {
	var p = new(request.ShopUserParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("商城注册请求参数有误", zap.Error(err))
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
	err := userService.Register(p)
	if err != nil {
		global.GA_LOG.Error("用户注册失败", zap.Error(err))
		response.ResponseError(c, config.CodeUserExist)
		return
	}
	response.ResponseSuccess(c, "注册成功")
}

// UserInfo 用户信息
func (u *ShopUserApi) UserInfo(c *gin.Context) {

}

// EditUserInfo 编辑用户信息
func (u *ShopUserApi) EditUserInfo(c *gin.Context) {

}
