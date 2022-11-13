package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/common/response"
	"go-admin/pkg/jwt"
	commonRedis "go-admin/pkg/redis"
	"go-admin/service"
	"go.uber.org/zap"
)

var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

// JWTAuth 基于JWT的认证中间件
func SysadminJwt() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		//token是否为空
		if token == "" {
			response.ResponseError(c, config.CodeNeedLogin) //需要登录
			c.Abort()                                       //阻止后续处理函数
			return
		}

		// 校验token是否在黑名单
		if global.GA_CONFIG.ApplicationConfig.UserRedis {
			is_black, _ := commonRedis.GetUserTokenBlackList(token)
			if is_black {
				response.ResponseError(c, config.CodeInvalidToken)
				c.Abort()
				return
			}
		}
		//解析token
		j := jwt.JWT{SigningKey: []byte(global.GA_CONFIG.JwtConfig.SigningKey)}
		claims, err := j.ParseToken(token)
		if err != nil {
			global.GA_LOG.Error("解析token失败 ", zap.Error(err), zap.String("TOKEN", token))
			response.ResponseError(c, config.CodeNeedLogin) //token无效
			c.Abort()                                       //阻止后续处理函数
			return
		}

		//global.GA_LOG.Info("token到期时间", zap.Any("ExpiresAt", claims.ExpiresAt))
		c.Set("userid", claims.UUID) //跨中间件设置值
		c.Next()                     //继续处理后续函数
	}
}
