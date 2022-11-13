package shop

import (
	"go-admin/global"
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

func GetUserCartList(userId uint) (list []*ShopCartItem, total int, err error) {
	res := global.GA_DB.Where("user_id =?  and is_deleted = 0", userId).Find(&list)
	total = int(res.RowsAffected)
	err = res.Error
	return
}

func AddUserCart(cart ShopCartItem) error {
	return global.GA_DB.Save(&cart).Error
}
