package redis

import (
	"context"
	"go-shop-api/global"
	"time"
)

var ctx = context.Background()

// SetUserTokenBlackList 设置用户token黑名单
func SetUserTokenBlackList(token string) (err error) {
	// 检查是否存在token
	res, err := global.GA_REDIS.SMembers(ctx, token).Result()
	if len(res) == 0 {
		return global.GA_REDIS.SAdd(ctx, "user:blacklist", token).Err()
	}
	return nil
}

// GetUserTokenBlackList 获取用户token是否在黑名单中
func GetUserTokenBlackList(token string) (val bool, err error) {
	return global.GA_REDIS.SIsMember(ctx, "user:blacklist", token).Result()
}

// ClearTokenBlackList 清除黑名单
func ClearTokenBlackList() error {
	return global.GA_REDIS.Del(ctx, "user:blacklist").Err()
}

// GetUserToken 获取用户Token
func GetUserToken(key string) (token string, err error) {
	return global.GA_REDIS.Get(context.Background(), key).Result()
}

// SetUserToken 设置用户Token
func SetUserToken(key string, token string) (err error) {
	timer := time.Duration(global.GA_CONFIG.JwtConfig.ExpiresTime) * time.Second
	return global.GA_REDIS.Set(ctx, key, token, timer).Err()
}

// SetOrderCloserTime 设置订单过期时间
func SetOrderCloserTime(key, orderNo string) (err error) {
	timer := time.Duration(global.GA_CONFIG.OrderCloserTime) * time.Second
	return global.GA_REDIS.SetEX(ctx, key, orderNo, timer).Err()
}
