package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go-shop-api/config"
	"go-shop-api/global"
	comRequest "go-shop-api/model/common/request"
	"go-shop-api/model/common/response"
	"go-shop-api/model/system"
	"go-shop-api/model/system/request"
	userResponse "go-shop-api/model/system/response"
	"go-shop-api/pkg/jwt"
	commonRedis "go-shop-api/pkg/redis"
	"go-shop-api/utils"
	"go.uber.org/zap"
	"time"
)

type BaseApi struct {
}

type ParamMerchantBalance struct {
	Mchid     string `json:"Mchid" binding:"required"`
	Timestamp int64  `json:"Timestamp" binding:"required"`
	Sign      string `json:"Sign" binding:"required"`
}

// Register 注册用户
func (b *BaseApi) Register(c *gin.Context) {
	//1.获取注册请求参数结构体
	var p = new(request.Register)
	//2.校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("注册请求参数有误", zap.Error(err))
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
	if err := userService.Register(p); err != nil {
		global.GA_LOG.Error("用户注册失败", zap.Error(err))
		response.ResponseError(c, config.CodeMenuExist)
		return
	}
	//4.返回响应
	response.ResponseSuccess(c, "创建成功")
}

func (b *BaseApi) CronTest(c *gin.Context) {
	time.Sleep(time.Second * 5)
	response.ResponseSuccess(c, "测试脚本333333")
}

// Login 登录用户
func (b *BaseApi) Login(c *gin.Context) {

	//1.校验登录参数
	var p = new(request.Login)

	//2.校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("登录参数有误", zap.Error(err))
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
	user, err := userService.Login(p)
	if err != nil {
		global.GA_LOG.Error("用户登录失败", zap.Error(err))
		//用户是否存在
		if errors.Is(err, userResponse.ErrorUserNotExit) {
			response.ResponseError(c, config.CodeUserNotExist)
			return
		}
		//用户密码错误
		response.ResponseError(c, config.CodeInvalidPassword)
		return
	}

	//4.生成token
	token, err := tokenNext(user)
	if err != nil {
		global.GA_LOG.Error("issue token err:", zap.Error(err))
		//系统繁忙
		response.ResponseError(c, config.CodeServerBusy)
	}

	//5.redis 存储 token
	if global.GA_CONFIG.ApplicationConfig.UserRedis {
		//如果旧token未自动生效删除旧token后存储token
		key := "user:sysadmin:token:" + p.Username
		userToken, err := commonRedis.GetUserToken(key)
		if err == redis.Nil {
			//写入用户token
			err = commonRedis.SetUserToken(key, token)
			if err != nil {
				global.GA_LOG.Error("set redis token err:", zap.Error(err))
				//系统繁忙
				response.ResponseError(c, config.CodeServerBusy)
			}

		} else if err != nil {
			global.GA_LOG.Error("get redis token err:", zap.Error(err))
			//系统繁忙
			response.ResponseError(c, config.CodeServerBusy)
		} else {
			// 将旧token写入黑名单
			if userToken != "" {
				err = commonRedis.SetUserTokenBlackList(userToken)
				if err != nil {
					global.GA_LOG.Error("old_token set blacklist err:", zap.Error(err))
					//系统繁忙
					response.ResponseError(c, config.CodeServerBusy)
				}
				// 重新写入token
				err = commonRedis.SetUserToken(key, token)
				if err != nil {
					global.GA_LOG.Error("set redis token err2:", zap.Error(err))
					//系统繁忙
					response.ResponseError(c, config.CodeServerBusy)
				}

			}
		}

	}

	//6.记录日志
	err = loginLogService.CreateLoginLog(c, p.Username, "1", "登陆成功")
	if err != nil {
		global.GA_LOG.Error("createLoginLog err: ", zap.Error(err))
		//系统繁忙
		response.ResponseError(c, config.CodeServerBusy)
	}

	//7.响应返回
	response.ResponseSuccess(c, gin.H{
		"token":  token,
		"userid": user.ID,
	})

}

//tokenNext 签发token
func tokenNext(user *system.SysUser) (token string, err error) {
	j := jwt.JWT{SigningKey: []byte(global.GA_CONFIG.JwtConfig.SigningKey)} // 唯一签名
	claims := j.CreateClaims(jwt.BaseClaims{
		UUID:        user.UUID,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err = j.CreateToken(claims)
	return
}

// UserInfo 用户信息
func (b *BaseApi) UserInfo(c *gin.Context) {

	//1.获取用户id
	jwtRes, err := utils.GetClaims(c)
	if err != nil {
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	//2.业务处理
	mp, err := userService.GetUserInfo(jwtRes.UUID)
	if err != nil {
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//3.响应返回
	response.ResponseSuccess(c, mp)
}

// Logout 退出登录
func (b *BaseApi) Logout(c *gin.Context) {
	//1将token写入黑名单
	token := c.Request.Header.Get("x-token")
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

// GetUserList 用户列表
func (b *BaseApi) GetUserList(c *gin.Context) {
	//1.请求结构体
	var p = new(request.GetUserList)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("用户列表请求参数有误", zap.Error(err))
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

	//2.查询用户
	list, total, err := userService.GetUserList(p)
	if err != nil {
		global.GA_LOG.Error("获取用户列表失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}

	//3.返回
	response.ResponseSuccess(c, response.PageResult{
		List:  list,
		Total: total,
		Page:  int(p.Page),
		Limit: int(p.Limit),
	})
}

// DelUser 删除用户
func (b *BaseApi) DelUser(c *gin.Context) {
	var p = new(comRequest.GetByIds)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("删除用户参数有误", zap.Error(err))
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
	err := userService.DelUser(p.ID)
	if err != nil {
		global.GA_LOG.Error("删除用户失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "删除成功")

}

// UpdateUser 更新用户
func (b *BaseApi) UpdateUser(c *gin.Context) {
	var p = new(request.EditUserParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("更新用户参数有误", zap.Error(err))
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
	err := userService.UpdateUser(p)
	if err != nil {
		global.GA_LOG.Error("更新用户失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "更新成功")
}

// ResetPassword 重置密码
func (b *BaseApi) EditPassword(c *gin.Context) {
	var p = new(request.ChangePasswordStruct)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("重置密码参数有误", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	err := userService.EditPassword(p)
	if err != nil {
		global.GA_LOG.Error("重置用户密码失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "更新密码成功")
}

// UpdateUserStatus 更新用户状态
func (b *BaseApi) UpdateUserStatus(c *gin.Context) {
	var p = new(request.EditUserStatus)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("更新用户状态", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	err := userService.EditUserStatus(p.UserId, p.Status)
	if err != nil {
		global.GA_LOG.Error("更新状态失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "更新状态成功")
}

// SetUserAuthority 用户设置角色
func (b *BaseApi) SetUserAuthority(c *gin.Context) {

}

// DelUserAvater 删除用户头像
func (b *BaseApi) DelUserAvater(c *gin.Context) {
	var p = new(request.DelUserAvaterParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("删除用户头像失败", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	if err := userService.DelUserAvater(p); err != nil {
		global.GA_LOG.Error("删除用户头像失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "头像删除成功")
}
