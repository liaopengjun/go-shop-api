package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type UserRouter struct {
}

// InitUserRouter 用户模块路由
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("register", baseApi.Register)                 // 注册账号
		userRouter.POST("logout", baseApi.Logout)                     // 用户退出
		userRouter.GET("userinfo", baseApi.UserInfo)                  // 获取用户信息
		userRouter.POST("userList", baseApi.GetUserList)              // 获取用户列表
		userRouter.POST("delUser", baseApi.DelUser)                   // 删除用户
		userRouter.POST("updateUser", baseApi.UpdateUser)             // 更新用户
		userRouter.POST("editPassword", baseApi.EditPassword)         // 更新密码
		userRouter.POST("setUserAuthority", baseApi.SetUserAuthority) // 设置用户角色
		userRouter.POST("updateUserStatus", baseApi.UpdateUserStatus) //更新用户状态
		userRouter.POST("delUserAvater", baseApi.DelUserAvater)       //更新用户状态
	}
}
