package shop

import (
	"go-shop-api/global"
	"time"
)

type ShopUserAddress struct {
	AddressId     int       `json:"addressId" form:"addressId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId        int       `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	UserName      string    `json:"userName" form:"userName" gorm:"column:user_name;comment:收货人姓名;type:varchar(30);"`
	UserPhone     string    `json:"userPhone" form:"userPhone" gorm:"column:user_phone;comment:收货人手机号;type:varchar(11);"`
	DefaultFlag   int       `json:"defaultFlag" form:"defaultFlag" gorm:"column:default_flag;comment:是否为默认 0-非默认 1-是默认;type:tinyint"`
	ProvinceName  string    `json:"provinceName" form:"provinceName" gorm:"column:province_name;comment:省;type:varchar(32);"`
	CityName      string    `json:"cityName" form:"cityName" gorm:"column:city_name;comment:城;type:varchar(32);"`
	RegionName    string    `json:"regionName" form:"regionName" gorm:"column:region_name;comment:区;type:varchar(32);"`
	DetailAddress string    `json:"detailAddress" form:"detailAddress" gorm:"column:detail_address;comment:收件详细地址(街道/楼宇/单元);type:varchar(64);"`
	IsDeleted     int       `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime    time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:添加时间;type:datetime"`
	UpdateTime    time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"`
}

func GetUserAddressInfo(userId uint, address_id int, default_flag int) (u *ShopUserAddress, err error) {
	db := global.GA_DB.Model(&ShopUserAddress{})
	if userId > 0 {
		db.Where("user_id=? and is_deleted = 0", userId)
	}
	if address_id > 0 {
		db.Where("address_id=?", address_id)
	}
	if default_flag > 0 {
		db.Where("default_flag = ?", default_flag)
	}
	err = db.First(&u).Error
	return
}
func GetUserAddressList(userId uint) (u []ShopUserAddress, err error) {
	err = global.GA_DB.Where(" user_id = ? and is_deleted = 0", userId).Find(&u).Error
	return
}

func SaveUserAddress(action string, u *ShopUserAddress) error {
	if action == "add" {
		return global.GA_DB.Create(&u).Error
	} else {
		return global.GA_DB.Updates(&u).Error
	}
}

func DelUserAddress(address_id int) (err error) {
	err = global.GA_DB.Where("address_id = ?", address_id).Delete(&ShopUserAddress{}).Error
	return
}
