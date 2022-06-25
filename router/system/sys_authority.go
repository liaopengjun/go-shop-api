package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type AuthorityRouter struct {
}

// InitUserRouter 用户模块路由
func (a *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority")
	var authorityApi = v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("addAuthority", authorityApi.CreateAuthority)      // 创建角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)   // 删除角色
		authorityRouter.POST("updateAuthority", authorityApi.UpdateAuthority)   // 更新角色
		authorityRouter.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
		authorityRouter.POST("setAuthority", authorityApi.SetAuthority)         // 设置角色权限
		authorityRouter.POST("getAuthInfo", authorityApi.GetAuthorityInfo)      // 角色详情
		authorityRouter.POST("authStatus", authorityApi.UpdateAuthorityStatus)  // 角色状态
		authorityRouter.POST("setAuthApi", authorityApi.SetAuthApi)             // 设置角色api
		authorityRouter.POST("getAuthApi", authorityApi.GetAuthApi)             //获取角色api列表

	}
}
