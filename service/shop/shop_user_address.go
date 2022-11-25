package shop

import (
	"go-admin/model/shop"
	"go-admin/model/shop/request"
	"time"
)

type ShopUserAddressService struct {
}

func (a *ShopUserAddressService) AddUserAddress(userId uint, userAddress *request.AddUserAddressParam) (err error) {

	//设置默认地址
	if userAddress.DefaultFlag == 1 {
		addressInfo, _ := shop.GetUserAddressInfo(userId, 0, 1)
		if addressInfo != nil {
			u2 := &shop.UserAddress{
				AddressId:   addressInfo.AddressId,
				DefaultFlag: 0,
				UpdateTime:  time.Time{},
			}
			err = shop.SaveUserAddress("edit", u2)
			if err != nil {
				return err
			}
		}
	}
	u := &shop.UserAddress{
		UserId:        int(userId),
		UserName:      userAddress.UserName,
		UserPhone:     userAddress.UserPhone,
		DefaultFlag:   userAddress.DefaultFlag,
		ProvinceName:  userAddress.ProvinceName,
		CityName:      userAddress.CityName,
		RegionName:    userAddress.RegionName,
		DetailAddress: userAddress.DetailAddress,
		IsDeleted:     0,
		CreateTime:    time.Time{},
		UpdateTime:    time.Time{},
	}
	return shop.SaveUserAddress("add", u)
}

func (a *ShopUserAddressService) DelUserAddress(address_id int) (err error) {
	return shop.DelUserAddress(address_id)
}

func (a *ShopUserAddressService) GetUserAddressInfo(userId uint, address_id int) (addressInfo *shop.UserAddress, err error) {
	addressInfo, err = shop.GetUserAddressInfo(userId, address_id, 0)
	return
}

func (a *ShopUserAddressService) GetUserAddressList(userId uint) (addressList []shop.UserAddress, err error) {
	addressList, err = shop.GetUserAddressList(userId)
	return
}

func (a *ShopUserAddressService) GetDefaultAddressInfo(userId uint) (addressInfo *shop.UserAddress, err error) {
	addressInfo, err = shop.GetUserAddressInfo(userId, 0, 1)
	return
}

func (a *ShopUserAddressService) EditUserAddress(userId uint, param *request.EditUserAddressParam) (err error) {

	//设置默认地址
	if param.DefaultFlag == 1 {
		addressInfo, _ := shop.GetUserAddressInfo(userId, 0, 1)
		if addressInfo != nil {
			u2 := &shop.UserAddress{
				AddressId:   addressInfo.AddressId,
				DefaultFlag: 0,
				UpdateTime:  time.Time{},
			}
			err = shop.SaveUserAddress("edit", u2)
			if err != nil {
				return err
			}
		}
	}

	u := &shop.UserAddress{
		AddressId:     param.AddressId,
		UserName:      param.UserName,
		UserPhone:     param.UserPhone,
		DefaultFlag:   param.DefaultFlag,
		ProvinceName:  param.ProvinceName,
		CityName:      param.CityName,
		RegionName:    param.RegionName,
		DetailAddress: param.DetailAddress,
		IsDeleted:     0,
		UpdateTime:    time.Time{},
	}
	return shop.SaveUserAddress("edit", u)
}
