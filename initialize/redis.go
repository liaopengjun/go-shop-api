package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-admin/global"
	"go.uber.org/zap"
)

//Redis 初始化redis
func Redis(){
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",global.GA_CONFIG.RedisConfig.Host, global.GA_CONFIG.RedisConfig.Port),
		Password:     global.GA_CONFIG.RedisConfig.Password,
		DB:           global.GA_CONFIG.RedisConfig.DB,
		PoolSize:     global.GA_CONFIG.RedisConfig.PoolSize,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		//global.GA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GA_REDIS = client
	}
	return
}

