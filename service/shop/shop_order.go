package shop

import (
	"go-shop-api/global"
	"go-shop-api/model/shop"
	"go-shop-api/model/shop/request"
	"go-shop-api/model/shop/response"
	"go-shop-api/utils"
)

type ShopOrderService struct {
}

func (o *ShopOrderService) CreateOrder(p *request.SaveOrderParam, userId uint) (orderCode string, err error) {
	UserCartItem, err := shop.GetCartItemDetailed(userId, p.CartItemIds)
	if err != nil {
		return "", response.ErrCartItem //购物车商品异常
	}

	var goodsIds []int64
	for _, item := range UserCartItem {
		goodsIds = append(goodsIds, item.GoodsId)
	}

	var goodsList []shop.ShopGoods
	err = global.GA_DB.Where("goods_id in ?", goodsIds).Find(&goodsList).Error
	if err != nil {
		return "", response.ErrGoodsNonExistent //商品不存在
	}

	//1.检查商品是否已下架
	var goodsStockNum = make(map[int]int)
	var isGoodsExist = 0
	var isErrOrderLowerShelf = 0
	for _, goodsInfo := range goodsList {
		goodsStockNum[goodsInfo.GoodsId] = goodsInfo.StockNum
		if goodsInfo.GoodsSellStatus == 1 {
			isErrOrderLowerShelf = 1
		}
		if utils.In_array(goodsInfo.GoodsId, goodsIds) {
			isGoodsExist = 1
		}
	}

	if isErrOrderLowerShelf == 1 {
		return "", response.ErrOrderLowerShelf //商品已下架
	}
	if isGoodsExist == 1 {
		return "", response.ErrGoodsNonExistent //商品不存在
	}

	//2.检查商品是否库存充足
	for _, cartItemInfo := range UserCartItem {
		if cartItemInfo.GoodsCount > goodsStockNum[int(cartItemInfo.GoodsId)] {
			return "", response.ErrGoodsInventory //商品库存不足
		}
	}

	//3.校验金额
	totalPrice := 0
	for _, itemInfo := range UserCartItem {
		totalPrice += itemInfo.SellingPrice * itemInfo.GoodsCount
	}
	if totalPrice <= 0 {
		return "", response.ErrGoodsTotalPrice //商品价格有误
	}

	//4.获取用户默认地址
	//userDefaultAddress, _ := shop.GetUserAddressInfo(userId, 0, 1)

	//事务开始
	//5.删除购物车数据
	//6.创建订单（快照）商品明细
	//事务结束

	return
}
