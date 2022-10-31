package initialize

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-admin/docs" //导入docs
	"go-admin/global"
	"go-admin/middleware"
	"go-admin/router"
	"net/http"
)

//Router 初始化路由
func Router() *gin.Engine {
	var Router = gin.Default()
	systemRouter := router.RouterGroupApp.System
	shopRouter := router.RouterGroupApp.Shop

	Router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler)) //api接口文档

	// Router.Use(middleware.Cors()) //跨域
	PrivateGroup := Router.Group("/api")
	PrivateGroup.StaticFS(global.GA_CONFIG.LocalConfig.Path, http.Dir(global.GA_CONFIG.LocalConfig.Path)) // 为用户头像和文件提供静态地址
	//基础路由部分
	systemRouter.InitBaseRouter(PrivateGroup)
	//中间件部分
	PrivateGroup.Use(middleware.JWTAuth(), middleware.Casbin())
	{
		systemRouter.InitMenuRouter(PrivateGroup)      //菜单路由
		systemRouter.InitUserRouter(PrivateGroup)      //用户路由
		systemRouter.InitAuthorityRouter(PrivateGroup) //权限路由
		systemRouter.InitApiRouter(PrivateGroup)       //系统api
		systemRouter.InitUploadRoute(PrivateGroup)     //上传文件
		systemRouter.InitLoginLogRouter(PrivateGroup)  //登陆日志
	}

	ShopGroup := Router.Group("/shop")
	shopRouter.InitShopBasicRouter(ShopGroup) //商城基础信息路由
	shopRouter.InitShopGoodsRouter(ShopGroup) //商城基础信息路由

	return Router
}
