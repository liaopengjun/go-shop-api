package shop

import (
	"errors"
	"go-admin/global"
	"go-admin/model/shop"
	"go-admin/model/shop/response"
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
	_, total, _ := shop.GetUserCartList(userId, 0)
	if total > 20 {
		return errors.New(" 购物车数量不超过20件")
	}
	// 检查是否商品已经加入购物车
	cart, itemCount, _ := shop.GetUserCartInfo(userId, goodsId)
	if itemCount > 0 {
		cartItemId := cart.CartItemId
		SumCount := cart.GoodsCount + count
		cartData := shop.ShopCartItem{
			CartItemId: cartItemId,
			GoodsCount: SumCount,
			UpdateTime: time.Now(),
		}
		return shop.UpdateCart(cartData)
	} else {
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

}

func (c *ShopCartService) GetCartList(userId uint, pageNumber int) (cartItemList []*response.CartItemResponse, total int, err error) {
	var goods_ids []int64
	var goodsList []shop.ShopGoods
	var list []*shop.ShopCartItem
	list, total, err = shop.GetUserCartList(userId, pageNumber)
	goodsMap := make(map[int64]shop.ShopGoods)

	for _, cartItem := range list {
		goods_ids = append(goods_ids, cartItem.GoodsId)
	}
	global.GA_DB.Where("goods_id in ?", goods_ids).Find(&goodsList)
	for _, goodsInfo := range goodsList {
		goodsMap[int64(goodsInfo.GoodsId)] = goodsInfo
	}
	for _, cartItem := range list {
		cartItemRes := response.CartItemResponse{
			CartItemId: cartItem.CartItemId,
			GoodsId:    cartItem.GoodsId,
			GoodsCount: cartItem.GoodsCount,
		}
		if _, ok := goodsMap[cartItem.GoodsId]; ok {
			goodsInfo := goodsMap[cartItem.GoodsId]
			cartItemRes.GoodsName = goodsInfo.GoodsName
			cartItemRes.GoodsCoverImg = goodsInfo.GoodsCoverImg
			cartItemRes.SellingPrice = goodsInfo.SellingPrice
		}
		cartItemList = append(cartItemList, &cartItemRes)
	}
	return
}

func (c *ShopCartService) GetCartAmout(userId uint) (total int64, err error) {
	return shop.GetUserCartCount(userId)
}
