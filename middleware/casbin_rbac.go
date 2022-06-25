package middleware

import (
	"github.com/gin-gonic/gin"
	"go-admin/config"
	"go-admin/model/common/response"
	"go-admin/service"
	"go-admin/utils"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 权限拦截器
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		waitUse, _ := utils.GetClaims(c)
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		// 判断策略中是否存在
		e := casbinService.Casbin()
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			response.ResponseError(c, config.CodeNoPermission)
			c.Abort()
			return
		}
	}
}
