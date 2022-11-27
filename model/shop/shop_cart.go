package shop

import (
	"go-shop-api/global"
	"time"
)

type ShopCartItem struct {
	CartItemId int       `json:"cartItemId" form:"cartItemId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId     uint      `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	GoodsId    int64     `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:关联商品id;type:bigint"`
	GoodsCount int       `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;comment:数量(最大为5);type:int"`
	IsDeleted  int       `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
	UpdateTime time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改时间;type:datetime"`
}

type UserCartItems struct {
	CartItemId    int
	GoodsId       int64
	GoodsName     string
	GoodsCount    int
	GoodsCoverImg string
	SellingPrice  int
}

func GetUserCartList(userId uint, pageNumber int) (list []*ShopCartItem, total int, err error) {
	if pageNumber > 0 {
		limit := 10
		offset := 10 * (pageNumber - 1)
		res := global.GA_DB.Where("user_id =?  and is_deleted = 0", userId).Limit(limit).Offset(offset).Find(&list)
		total = int(res.RowsAffected)
		err = res.Error
	} else {
		res := global.GA_DB.Where("user_id =?  and is_deleted = 0", userId).Find(&list)
		total = int(res.RowsAffected)
		err = res.Error
	}
	return
}

func GetUserCartInfo(userId uint, stockId int64) (cart *ShopCartItem, cartCount int64, err error) {
	res := global.GA_DB.Where("user_id = ? and is_deleted = 0 and goods_id = ?", userId, stockId).First(&cart)
	cartCount = res.RowsAffected
	err = res.Error
	return
}

func AddUserCart(cart ShopCartItem) error {
	return global.GA_DB.Save(&cart).Error
}

func UpdateCart(cart ShopCartItem) error {
	return global.GA_DB.Updates(cart).Error
}
func DelCart(cartItemID int) error {
	return global.GA_DB.Delete(&ShopCartItem{}, cartItemID).Error
}

func GetUserCartCount(userId uint) (total int64, err error) {
	err = global.GA_DB.Model(&ShopCartItem{}).Where("user_id =?  and is_deleted = 0", userId).Count(&total).Error
	return
}

func GetCartItemDetailed(userId uint, cartItemIds []int) (UserCartItems []UserCartItems, err error) {
	err = global.GA_DB.Model(&ShopCartItem{}).Select("shop_cart_items.cart_item_id,shop_cart_items.goods_id,shop_cart_items.goods_count,shop_goods.selling_price,shop_goods.goods_name,shop_goods.goods_cover_img").Joins("left join shop_goods on shop_cart_items.goods_id = shop_goods.goods_id").Where("shop_cart_items.user_id = ? and shop_cart_items.is_deleted =0 and shop_cart_items.cart_item_id in (?) ", userId, cartItemIds).Scan(&UserCartItems).Error
	return
}
