package middleware

import (
	"github.com/gin-gonic/gin"
	"go-shop-api/config"
	"go-shop-api/global"
	"go-shop-api/model/common/response"
	"go-shop-api/pkg/jwt"
	commonRedis "go-shop-api/pkg/redis"
	"go.uber.org/zap"
)

// JWT 基于JWT的认证中间件
func ShopJwt() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
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
		c.Set("shop_userid", claims.ID) //跨中间件设置值
		c.Next()                        //继续处理后续函数
	}
}
