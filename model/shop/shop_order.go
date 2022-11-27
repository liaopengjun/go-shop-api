package shop

import "time"

type ShopOrder struct {
	OrderId     int        `json:"orderId" form:"orderId" gorm:"primarykey;AUTO_INCREMENT"`
	OrderNo     string     `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:订单号;type:varchar(20);"`
	UserId      int        `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	TotalPrice  int        `json:"totalPrice" form:"totalPrice" gorm:"column:total_price;comment:订单总价;type:int"`
	PayStatus   int        `json:"payStatus" form:"payStatus" gorm:"column:pay_status;comment:支付状态:0.未支付,1.支付成功,-1:支付失败;type:tinyint"`
	PayType     int        `json:"payType" form:"payType" gorm:"column:pay_type;comment:0.无 1.支付宝支付 2.微信支付;type:tinyint"`
	PayTime     *time.Time `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;type:datetime"`
	OrderStatus int        `json:"orderStatus" form:"orderStatus" gorm:"column:order_status;comment:订单状态:0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭;type:tinyint"`
	ExtraInfo   string     `json:"extraInfo" form:"extraInfo" gorm:"column:extra_info;comment:订单body;type:varchar(100);"`
	UserName    string     `json:"user_name" form:"user_name" gorm:"column:user_name;comment:收货人姓名;type:varchar(30);"`
	UserPhone   string     `json:"user_phone" form:"user_phone" gorm:"column:user_phone;comment:收货人手机号;type:varchar(11);"`
	UserAddress string     `json:"user_address" form:"user_address" gorm:"column:user_address;comment:收货人地址;type:varchar(100);"`
	IsDeleted   int        `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime  time.Time  `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
	UpdateTime  time.Time  `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改时间;type:datetime"`
}

type ShopOrderItem struct {
	OrderItemId   int       `json:"orderItemId" gorm:"primarykey;AUTO_INCREMENT"`
	OrderId       int       `json:"orderId" form:"orderId" gorm:"column:order_id;;type:bigint"`
	GoodsId       int       `json:"goodsId" form:"goodsId" gorm:"column:goods_id;;type:bigint"`
	GoodsName     string    `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名;type:varchar(200);"`
	GoodsCoverImg string    `json:"goodsCoverImg" form:"goodsCoverImg" gorm:"column:goods_cover_img;comment:商品主图;type:varchar(200);"`
	SellingPrice  int       `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;comment:商品实际售价;type:int"`
	GoodsCount    int       `json:"goodsCount" form:"goodsCount" gorm:"column:goods_count;;type:bigint"`
	CreateTime    time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
}
