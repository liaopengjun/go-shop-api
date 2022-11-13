package shop

import (
	"errors"
	"go-admin/model/shop"
	"time"
)

type ShopCartService struct {
}

func (c *ShopCartService) AddCart(userId uint, goodsId int64, count int) error {
	//商品是否为空
	_, err := shop.GetGoodsDetail(goodsId)
	if err != nil {
		return errors.New(" 商品为空")
	}
	//购物车数量不超过20件商品
	_, total, _ := shop.GetUserCartList(userId)
	if total > 20 {
		return errors.New(" 购物车数量不超过20件")
	}
	//插入数据
	cartData := shop.ShopCartItem{
		UserId:     userId,
		GoodsId:    goodsId,
		GoodsCount: count,
		IsDeleted:  0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	return shop.AddUserCart(cartData)
}
