package main

import (
	"context"
	"fmt"
	"go-shop-api/global"
	"go-shop-api/initialize"
	"go-shop-api/model/shop"
	"strings"
	"time"
)

var ctx = context.Background()

func main() {
	global.GA_VP = initialize.Viper("./config/config.yaml")
	global.GA_DB = initialize.Gorm() // 初始化数据库
	initialize.Redis()               //初始化redis
	pubsub := global.GA_REDIS.Subscribe(ctx, "__keyevent@0__:expired")
	defer pubsub.Close()
	fmt.Println("过期订单监听中...")
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		orderNo := strings.Split(msg.Payload, ":")[3]
		err = global.GA_DB.Model(shop.ShopOrder{}).Where("is_deleted = 0 and order_status = 0 and pay_status = 0 and order_no = ?", orderNo).Update("order_status", -2).Error
		if err != nil {
			panic(err)
		}
		fmt.Println(orderNo + "---超时关闭订单成功")
		time.Sleep(time.Millisecond * 100)
	}
}
