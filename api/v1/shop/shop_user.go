package shop

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/model/shop"
	"go-admin/model/shop/request"
	shopResponse "go-admin/model/shop/response"
	"go-admin/pkg/jwt"
	commonRedis "go-admin/pkg/redis"
	"go-admin/utils"
	"go.uber.org/zap"
)

type ShopUserApi struct {
}

// Login 用户登录
func (u *ShopUserApi) Login(c *gin.Context) {
	var p = new(request.ShopUserParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("商城登录请求参数有误", zap.Error(err))
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

	user, err := userService.Login(p)
	if err != nil {
		global.GA_SHOPLOG.Error("用户登录失败", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidPassword)
		return
	}

	token, err := tokenNext(user)
	if err != nil {
		global.GA_SHOPLOG.Error("issue shopuser token err:", zap.Error(err))
		//系统繁忙
		response.ResponseError(c, config.CodeServerBusy)
	}

	if global.GA_CONFIG.ApplicationConfig.UserRedis {
		//如果旧token未自动生效删除旧token后存储token
		key := "user:shop:token:" + user.LoginName
		userToken, err := commonRedis.GetUserToken(key)
		if err == redis.Nil {
			//写入用户token
			err = commonRedis.SetUserToken(key, token)
			if err != nil {
				global.GA_SHOPLOG.Error("set redis token err:", zap.Error(err))
				//系统繁忙
				response.ResponseError(c, config.CodeServerBusy)
			}

		} else if err != nil {
			global.GA_SHOPLOG.Error("get redis token err:", zap.Error(err))
			//系统繁忙
			response.ResponseError(c, config.CodeServerBusy)
		} else {
			// 将旧token写入黑名单
			if userToken != "" {
				err = commonRedis.SetUserTokenBlackList(userToken)
				if err != nil {
					global.GA_SHOPLOG.Error("old_token set blacklist err:", zap.Error(err))
					//系统繁忙
					response.ResponseError(c, config.CodeServerBusy)
				}
				// 重新写入token
				err = commonRedis.SetUserToken(key, token)
				if err != nil {
					global.GA_SHOPLOG.Error("set redis token err2:", zap.Error(err))
					//系统繁忙
					response.ResponseError(c, config.CodeServerBusy)
				}

			}
		}

	}
	response.ResponseSuccess(c, token)
}

// OutLogin 退出登录
func (u *ShopUserApi) OutLogin(c *gin.Context) {
	//1将token写入黑名单
	token := c.Request.Header.Get("token")
	if global.GA_CONFIG.ApplicationConfig.UserRedis {
		err := commonRedis.SetUserTokenBlackList(token)
		if err != nil {
			response.ResponseError(c, config.CodeServerBusy)
			return
		}
	}
	//2.响应返回
	response.ResponseSuccess(c, "退出成功")

}

// Register 注册
func (u *ShopUserApi) Register(c *gin.Context) {
	var p = new(request.ShopUserParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("商城注册请求参数有误", zap.Error(err))
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
		global.GA_SHOPLOG.Error("用户注册失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "注册成功")
}

// UserInfo 用户信息
func (u *ShopUserApi) UserInfo(c *gin.Context) {
	shop_userid, ok := c.Get("shop_userid")
	if !ok {
		response.ResponseError(c, config.CodeNeedLogin)
		return
	}
	user, err := userService.GetUserInfo(shop_userid.(uint))
	if err != nil {
		global.GA_SHOPLOG.Error("获取用户失败", zap.Error(err))
		response.ResponseError(c, config.CodeUserExist)
		return
	}
	response.ResponseSuccess(c, user)
}

// EditUserInfo 编辑用户信息
func (u *ShopUserApi) EditUserInfo(c *gin.Context) {
	var p = new(request.ShopEditUserParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_SHOPLOG.Error("商城编辑用户请求参数有误", zap.Error(err))
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

	err := userService.EditUser(p)
	if err != nil {
		global.GA_SHOPLOG.Error("编辑用户信息失败", zap.Error(err))
		if errors.Is(err, shopResponse.ErrorPasswordWrong) {
			response.ResponseError(c, config.CodeInvalidPassword)
			return
		}
		response.ResponseError(c, config.CodeServerBusy)
		return
	}

	response.ResponseSuccess(c, "编辑信息成功")
}

//tokenNext 签发token
func tokenNext(user *shop.ShopUser) (token string, err error) {
	j := jwt.JWT{SigningKey: []byte(global.GA_CONFIG.JwtConfig.SigningKey)} // 唯一签名
	claims := j.CreateClaims(jwt.BaseClaims{
		UUID:        uuid.UUID{},
		Username:    user.LoginName,
		ID:          uint(user.UserId),
		AuthorityId: "0",
	})
	return j.CreateToken(claims)
}
