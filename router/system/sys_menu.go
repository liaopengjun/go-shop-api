package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-admin/api/v1"
)

type MenuRouter struct {
}

// InitMenuRouter 菜单模块路由
func (s *UserRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu")
	var MenuApi = v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		menuRouter.POST("addMenu", MenuApi.AddMenu)                   // 新增菜单
		menuRouter.POST("deleteMenu", MenuApi.DeleteMenu)             // 删除菜单
		menuRouter.POST("updateMenu", MenuApi.UpdateMenu)             // 更新菜单
		menuRouter.POST("addMenuAuthority", MenuApi.AddMenuAuthority) // 绑定角色关系

		menuRouter.POST("getMenuTreeList", MenuApi.GetMenuTree) // 获取菜单树
		menuRouter.POST("getMenuList", MenuApi.GetMenuList)     // menu列表
		menuRouter.POST("getMenuInfo", MenuApi.GetMenuInfo)     // menu详情

		menuRouter.POST("getAuthorityMenuList", MenuApi.GetAuthorityMenuList) //角色菜单
		menuRouter.POST("getUserMenuList", MenuApi.GetUserMenuList)           //侧边栏菜单

	}
}
