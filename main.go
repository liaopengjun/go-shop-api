package main

import (
	"fmt"
	"go-admin/api/v1/system"
	"go-admin/global"
	"go-admin/initialize"
)

// @title lpjcode
// @version 1.0
// @description go-admin
// @termsOfService http://swagger.io/terms/
// @contact.name
// @contact.url http://www.swagger.io/support
// @contact.email 1337404942@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 这里写接口服务的host
// @BasePath 这里写base path

func main() {
	global.GA_VP = initialize.Viper() // 初始化Viper
	global.GA_LOG = initialize.Zap()  // 初始化zap日志库
	global.GA_DB = initialize.Gorm()  // 初始化数据库
	system.Trans("zh")                //gin框架内置校验翻译器
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
