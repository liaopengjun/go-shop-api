package main

import (
	"fmt"
	"go-shop-api/api/v1/system"
	"go-shop-api/global"
	"go-shop-api/initialize"
)

func main() {
	global.GA_VP = initialize.Viper()                                   // 初始化Viper
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

//
//APPID
//2016092800619304
//
//应用名称
//sandbox应用:2088102177717881
//
//绑定的商家账号（PID）
//2088102177717881
//
//应用公钥
//MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6W7qHtGP/ZYHAaYVY9LSTIKS+ns6DvwJ1WhgDKiknTDTru8zbyaiOn7seR8VaP3lDtHT6MQywyq3XuYLsks11RaSodRmtPLCHvEYrBxgnmmyUeH0f3jycmDs4kcZHInar2QEeJHom686QtvmcDH5u0CVFTfCFcpoFb4zXSQeDATbYEXHk53j/DE3f0T03d83WkmBFyM91sgecoF3OqdvF8j6CcUkeIsTCq4hgSHWT1bi5sWD9r+twBGwE9j24DxH9hHK9KPqPxa7LHRgwpjX2qZj18lqYKJ6i3PCwTkSZWOKG0M7pvtNVdLfcNHVAwhJ4w5m+4ip8Ky2jAbsWj6FbQIDAQAB
//
//私钥
//MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDpbuoe0Y/9lgcBphVj0tJMgpL6ezoO/AnVaGAMqKSdMNOu7zNvJqI6fux5HxVo/eUO0dPoxDLDKrde5guySzXVFpKh1Ga08sIe8RisHGCeabJR4fR/ePJyYOziRxkcidqvZAR4keibrzpC2+ZwMfm7QJUVN8IVymgVvjNdJB4MBNtgRceTneP8MTd/RPTd3zdaSYEXIz3WyB5ygXc6p28XyPoJxSR4ixMKriGBIdZPVuLmxYP2v63AEbAT2PbgPEf2Ecr0o+o/FrssdGDCmNfapmPXyWpgonqLc8LBORJlY4obQzum+01V0t9w0dUDCEnjDmb7iKnwrLaMBuxaPoVtAgMBAAECggEAKtIT1G9w0H3S7zR+O/+SYEKbn6M3NUR6sVEiPXA6MjOgwThT4RHfNIfP7TAMh2P7vsoy60ICZdbSKHBeuOgfCeXIJDOIW60kevSTKA9UkfqloWunpDKVlvG8wo10R83p9b6NK2jomJZ+Q4F/NhmUzAq+zlOaINEuYr6vQWi8sktgFEY9H1nJ3Y/eD6MeR9vx8WJkzgAqwdea2v/XH6+bHPzdKGSpLU4J6Nxrao7ceoyhdZZBJEIN//H/ZMXTCLJfxrdCGFbu73Zw0ej6VkGpFoFoTg3ymQI9s+HZv5M2R2k4RVZdCK//HAbTGfcqCozj7A/K9AxAwK3KjHFW2VKCQQKBgQD6afBPPhGjEf17I1SZiL9g6jpD92+bt0nGT292M2uppqmQF8olWxAygRZvxzyhjI/YRuTHbWJQI8xvuwyNFESD7tPbWl30a/iYiWGFACcnjIfjidw3StiU1B7HF+vmh/Qpy5G2vxfcApejZ21MNpoCcOiIPqlbZVBvY3FGEOVJeQKBgQDupADcAA2cnzWTOqbvkAUC1E5O9cYZwREh3cONn46smZ/8Vok1gGMsByL+4Dgan779+JiTbXx+EByM/WuHQcfNCKr88QEcJV6bwGEA0/K0F414REDqMhebOeAgucNFycUSPAvtRJ/b4Zs/bUzFEUbU/XYV2c6NbYlwNX3FvtJSlQKBgGCsh3bAOqTe7CIe2KlRbrjmlEnq+659C0FBJ4HVhin/ypRzaroTNuSYi2Xp4BFqJ5pSfD41j2/q8iDscIMCoRTiHe4gLAeRq26QExL6pSMSkN+aOGcsQQLsBVnNdWgRcoS0L3QCwB5S7eHKqxpyNfDdUBhRQKalLXFjTbbEDRDZAoGAbk5mi0qHAC9jX0OMKE3E0zL5Y2wdfogMeD/+hTcMhuGX3tbNI2rN7Gr8FR3lMQFIEjLXq8W+9rJR0CXPjzyrsy1fg/2OiskHOy0oaW6O0AnW/ZFnBBnVaY3N+LKE/XwvWKdix/Chh8x3q1DFXI4I1Ki37Y+49wx7q9893KrAoyUCgYEAyCpi1FpSMLu1X2lL11xfXO5iA/3hKiXcIL1XzW+EnwnCrxK60eNqn/A1Yjegf1oyRxq5Qnc0AGE1f03k9O5KjmaYajeZwH4CRAmL+bQ4ANCQx13fBzjvVQD03/aSgi897M99sqy3npXyJ1/ONimXBmUBbV3v5CgqjoWvMqiXaDE=
//
//支付宝公钥
//MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6bZwQJiLBN/4EB5jcBw9vilwX1AZ8N7oo+7G30VXgmNEr1Bp2hq7tys8e6E3sNt+6XrUoRCCtI5ljv3jRpmWEhltR+gu47mmfU7eV5UvB9rsTMjNxZun1A9b3ubwTXRglbrw+kJhFZsaf0gIbyqmrtLgjmJ76ZVPvhLWCqARCE6Nlj2qT6OOSBTVIcyBIPTOVgcAa96Qiib0qAWrbgDk0QGrfmmQM1+xe+MyKVPBHmWOtYeXYnLJ1dDKP9jFkt6GgmChX34RTH4RehAI1JCMX63oa0/X6zsy+KIc6QD8o0WzedaVrCJ4NPqOOfk7Ao5hAf1WHYkEBQtUnPYmPgYVvwIDAQAB
