package main

import (
	"fmt"
	"go-shop-api/api/v1/system"
	"go-shop-api/global"
	"go-shop-api/initialize"
)

func main() {
	global.GA_VP = initialize.Viper("")                                 // 初始化Viper
	global.GA_LOG = initialize.Zap(global.GA_CONFIG.LogConfig.Director) // 初始化zap日志库
	global.GA_SHOPLOG = initialize.Zap(global.GA_CONFIG.LogConfig.ShopDirector)
	global.GA_DB = initialize.Gorm() // 初始化数据库
	system.Trans("zh")               //gin框架内置校验翻译器
	if global.GA_DB != nil {
		//自动迁移文件
		initialize.RegisterTables(global.GA_DB)
		db, _ := global.GA_DB.DB()
		//释放资源
		defer db.Close()
	}
	if global.GA_CONFIG.ApplicationConfig.UserRedis {
		initialize.Redis()            //初始化redis
		defer global.GA_REDIS.Close() //释放资源
	}
	Router := initialize.Router()
	Router.Run(fmt.Sprintf(":%d", global.GA_CONFIG.ApplicationConfig.Port))
}
