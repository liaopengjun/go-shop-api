package shop

import (
	"errors"
	"go-shop-api/global"
	"go-shop-api/model/common/enum"
	"go-shop-api/model/shop"
	"go-shop-api/model/shop/request"
	"go-shop-api/model/shop/response"
	"go-shop-api/pkg/pay"
	"go-shop-api/pkg/redis"
	"go-shop-api/utils"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ShopOrderService struct {
}

func (o *ShopOrderService) CreateOrder(p *request.SaveOrderParam, userId uint) (orderCode string, err error) {

	//查询购物车明细
	UserCartItem, err := shop.GetCartItemDetailed(userId, p.CartItemIds)
	if err != nil {
		return "", response.ErrCartItem //购物车商品异常
	}

	var goodsIds []int
	var itemIdList []int
	for _, item := range UserCartItem {
		goodsIds = append(goodsIds, int(item.GoodsId))
		itemIdList = append(itemIdList, item.CartItemId)
	}

	// 查询商品明细
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
		if !utils.In_array(goodsInfo.GoodsId, goodsIds) {
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
		totalPrice += totalPrice + itemInfo.SellingPrice*itemInfo.GoodsCount
	}
	if totalPrice <= 0 {
		return "", response.ErrGoodsTotalPrice //商品价格有误
	}

	//4.获取用户默认地址
	userDefaultAddress, _ := shop.GetUserAddressInfo(userId, p.AddressId, 0)
	orderCode = utils.GenOrderNo()

	err = global.GA_DB.Transaction(func(tx *gorm.DB) error {
		//事务开始
		//5.删除购物车数据
		TxErr := tx.Where("cart_item_id in ?", itemIdList).Updates(shop.ShopCartItem{IsDeleted: 1}).Error
		if TxErr != nil {
			return TxErr
		}

		closerTime := time.Now().Unix() + global.GA_CONFIG.OrderCloserTime
		timeStr := time.Unix(closerTime, 0).Format(global.TIME_FORMAT)
		t, _ := time.ParseInLocation(global.TIME_FORMAT, timeStr, time.Local)

		//创建订单
		orderData := shop.ShopOrder{
			OrderNo:     orderCode,
			UserId:      int(userId),
			TotalPrice:  totalPrice,
			ExtraInfo:   "",
			UserName:    userDefaultAddress.UserName,
			UserPhone:   userDefaultAddress.UserPhone,
			UserAddress: userDefaultAddress.DetailAddress,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
			OverdueTime: t,
		}
		if TxErr = tx.Save(&orderData).Error; TxErr != nil {
			return TxErr
		}
		orderId := orderData.OrderId
		//6.商品明细
		var shopOrderItem []shop.ShopOrderItem
		for _, cartItem := range UserCartItem {
			orderItem := shop.ShopOrderItem{
				OrderId:       orderId,
				GoodsId:       int(cartItem.GoodsId),
				GoodsName:     cartItem.GoodsName,
				GoodsCoverImg: cartItem.GoodsCoverImg,
				SellingPrice:  cartItem.SellingPrice,
				GoodsCount:    cartItem.GoodsCount,
				CreateTime:    time.Now(),
			}
			shopOrderItem = append(shopOrderItem, orderItem)
		}
		if TxErr = tx.Save(&shopOrderItem).Error; TxErr != nil {
			return TxErr
		}
		//事务结束
		return nil
	})
	//写入redis
	key := "shop:order:closer:" + orderCode
	err = redis.SetOrderCloserTime(key, orderCode)
	if err != nil {
		return orderCode, response.ErrCreateOrder
	}

	return orderCode, err
}

func (o *ShopOrderService) OrderPay2(p *request.OrderPayParam) (result map[string]interface{}, err error) {
	//检查订单有效状态
	orderInfo, err := shop.GetOrderInfo(p.OrderNo)
	if err != nil || orderInfo == nil {
		return nil, errors.New("订单信息有误")
	}
	var payParam map[string]interface{}
	payParam["subject"] = "订单沙箱环境"
	payParam["out_trade_no"] = orderInfo.OrderNo
	payParam["total_amount"] = orderInfo.TotalPrice
	payParam["notify_url"] = "111"
	payService := pay.NewPay(p.PayType)
	return payService.Pay("WapPay", payParam)
}

func (o *ShopOrderService) OrderPay(p *request.OrderPayParam) (err error) {
	//检查订单有效状态
	orderInfo, err := shop.GetOrderInfo(p.OrderNo)
	if err != nil || orderInfo == nil {
		return errors.New("订单信息有误")
	}

	PayType, _ := strconv.Atoi(p.PayType)
	now := time.Now()
	orderData := &shop.ShopOrder{
		OrderId:     orderInfo.OrderId,
		OrderNo:     p.OrderNo,
		PayStatus:   1,
		PayType:     PayType,
		PayTime:     now,
		OrderStatus: 1,
		UpdateTime:  time.Now(),
	}
	err = shop.SaveOrder(orderData)
	return
}

func (o *ShopOrderService) GetOrderList(p *request.OrderListParam, userId uint) (orderList []response.OrderResponse, total int64, err error) {
	orderRes, total, err := shop.GetOrderList(userId, p.PageNumber, p.Status)
	if total > 0 {

		//订单id
		var orderIds []int
		for _, orderInfo := range orderRes {
			orderIds = append(orderIds, orderInfo.OrderId)
		}

		//获取订单的商品信息
		itemList, _ := shop.GetOrderItemList(orderIds)
		var orderItemMap = make(map[int][]response.OrderItemResponse)
		for _, orderItemInfo := range itemList {
			itemData := response.OrderItemResponse{
				GoodsId:       orderItemInfo.GoodsId,
				GoodsName:     orderItemInfo.GoodsName,
				GoodsCount:    orderItemInfo.GoodsCount,
				GoodsCoverImg: orderItemInfo.GoodsCoverImg,
				SellingPrice:  orderItemInfo.SellingPrice,
			}
			orderItemMap[orderItemInfo.OrderId] = append(orderItemMap[orderItemInfo.OrderId], itemData)
		}

		//组合订单信息
		for _, orderInfo := range orderRes {
			orderItemSli := orderItemMap[orderInfo.OrderId]
			_, statusTxt := enum.GetNewBeeMallOrderStatusEnumByStatus(orderInfo.OrderStatus)
			orderData := response.OrderResponse{
				OrderId:           orderInfo.OrderId,
				OrderNo:           orderInfo.OrderNo,
				TotalPrice:        orderInfo.TotalPrice,
				PayType:           orderInfo.PayType,
				OrderStatus:       orderInfo.OrderStatus,
				OrderStatusString: statusTxt,
				CreateTime:        orderInfo.CreateTime,
				OrderItemResponse: orderItemSli,
			}
			orderList = append(orderList, orderData)
		}

	}
	return orderList, total, err
}

func (o *ShopOrderService) GetOrderDetail(orderNo string) (orderRes response.OrderDetailResponse, err error) {

	orderInfo, err := shop.GetOrderInfo(orderNo)
	if err != nil || orderInfo == nil {
		return orderRes, errors.New("订单信息有误")
	}

	itemList, err := shop.GetOrderItemList([]int{orderInfo.OrderId})
	if err != nil || itemList == nil {
		return orderRes, errors.New("订单信息有误")
	}

	//组合订单信息
	_, statusTxt := enum.GetNewBeeMallOrderStatusEnumByStatus(orderInfo.OrderStatus)
	_, PayTypeString := enum.GetNewBeeMallOrderPayTypeEnumByStatus(orderInfo.PayType)

	orderRes = response.OrderDetailResponse{
		OrderNo:           orderInfo.OrderNo,
		TotalPrice:        orderInfo.TotalPrice,
		PayStatus:         orderInfo.PayStatus,
		PayType:           orderInfo.PayType,
		PayTypeString:     PayTypeString,
		PayTime:           orderInfo.PayTime,
		OrderStatus:       orderInfo.OrderStatus,
		OrderStatusString: statusTxt,
		CreateTime:        orderInfo.CreateTime,
	}
	var orderItemSli []response.OrderItemResponse
	for _, orderItemInfo := range itemList {
		orderItemData := response.OrderItemResponse{
			GoodsId:       orderItemInfo.GoodsId,
			GoodsName:     orderItemInfo.GoodsName,
			GoodsCount:    orderItemInfo.GoodsCount,
			GoodsCoverImg: orderItemInfo.GoodsCoverImg,
			SellingPrice:  orderItemInfo.SellingPrice,
		}
		orderItemSli = append(orderItemSli, orderItemData)
	}
	orderRes.OrderItemResponse = orderItemSli
	return
}

func (o *ShopOrderService) CancelOrder(orderNo string) error {
	//检查订单有效状态
	orderInfo, err := shop.GetOrderInfo(orderNo)
	if err != nil || orderInfo == nil {
		return errors.New("订单信息有误")
	}
	orderData := &shop.ShopOrder{
		OrderId:     orderInfo.OrderId,
		OrderStatus: -1, //手动关闭
		UpdateTime:  time.Now(),
	}
	return shop.SaveOrder(orderData)
}
func (o *ShopOrderService) ConfirmTheGoods(orderNo string) error {
	//检查订单有效状态
	orderInfo, err := shop.GetOrderInfo(orderNo)
	if err != nil || orderInfo == nil {
		return errors.New("订单信息有误")
	}
	orderData := &shop.ShopOrder{
		OrderId:     orderInfo.OrderId,
		OrderStatus: 4, //手动关闭
		UpdateTime:  time.Now(),
	}
	return shop.SaveOrder(orderData)
}
