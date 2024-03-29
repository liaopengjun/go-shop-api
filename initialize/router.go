package initialize

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-shop-api/docs" //导入docs
	"go-shop-api/global"
	"go-shop-api/middleware"
	"go-shop-api/router"
	"net/http"
)

//Router 初始化路由
func Router() *gin.Engine {
	var Router = gin.Default()
	systemRouter := router.RouterGroupApp.System
	shopRouter := router.RouterGroupApp.Shop

	Router.Use(middleware.Cors())                                     //跨域
	Router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler)) //api接口文档
	PrivateGroup := Router.Group("/api")
	PrivateGroup.StaticFS(global.GA_CONFIG.LocalConfig.Path, http.Dir(global.GA_CONFIG.LocalConfig.Path)) // 为用户头像和文件提供静态地址

	systemRouter.InitBaseRouter(PrivateGroup)
	PrivateGroup.Use(middleware.SysadminJwt(), middleware.Casbin())
	{
		systemRouter.InitMenuRouter(PrivateGroup)      //菜单路由
		systemRouter.InitUserRouter(PrivateGroup)      //用户路由
		systemRouter.InitAuthorityRouter(PrivateGroup) //权限路由
		systemRouter.InitApiRouter(PrivateGroup)       //系统api
		systemRouter.InitUploadRoute(PrivateGroup)     //上传文件
		systemRouter.InitLoginLogRouter(PrivateGroup)  //登陆日志
		systemRouter.InitJobRouterRouter(PrivateGroup) //计划任务
	}

	ShopGroup := Router.Group("/shop")
	shopRouter.InitShopBasicRouter(ShopGroup) //商城基础信息路由
	ShopGroup.Use(middleware.ShopJwt())
	{
		shopRouter.InitShopGoodsRouter(ShopGroup)       //商城商品模块路由
		shopRouter.InitShopUserRouter(ShopGroup)        //商城用户路由
		shopRouter.InitShopCategoryRouter(ShopGroup)    //商城商品分类路由
		shopRouter.InitShopOrderRouter(ShopGroup)       //商城订单路由
		shopRouter.InitShopCartRouter(ShopGroup)        //商城购物车路由
		shopRouter.InitShopUserAddressRouter(ShopGroup) //商城购物车路由
	}

	return Router
}
