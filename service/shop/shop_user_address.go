package shop

import (
	"errors"
	"github.com/jinzhu/copier"
	"go-shop-api/model/shop"
	"go-shop-api/model/shop/request"
	"time"
)

type ShopUserAddressService struct {
}

func (a *ShopUserAddressService) AddUserAddress(userId uint, userAddress *request.AddUserAddressParam) (err error) {

	//设置默认地址
	var defaultAddress shop.ShopUserAddress
	if userAddress.DefaultFlag == 1 {
		defaultAddress, _ = shop.GetUserAddressInfo(userId, 0, 1)
		if defaultAddress != (shop.ShopUserAddress{}) {
			defaultAddress.UpdateTime = time.Now()
			defaultAddress.DefaultFlag = 0
			err = shop.SaveUserAddress("edit", &defaultAddress)
			if err != nil {
				return err
			}
		}
	}

	var addUserAddress shop.ShopUserAddress
	copier.Copy(&addUserAddress, userAddress)
	addUserAddress.UserId = int(userId)
	addUserAddress.CreateTime = time.Now()
	addUserAddress.UpdateTime = time.Now()
	return shop.SaveUserAddress("add", &addUserAddress)
}

func (a *ShopUserAddressService) DelUserAddress(address_id int) (err error) {
	return shop.DelUserAddress(address_id)
}

func (a *ShopUserAddressService) GetUserAddressInfo(userId uint, address_id int) (addressInfo shop.ShopUserAddress, err error) {
	addressInfo, err = shop.GetUserAddressInfo(userId, address_id, 0)
	return
}

func (a *ShopUserAddressService) GetUserAddressList(userId uint) (addressList []shop.ShopUserAddress, err error) {
	addressList, err = shop.GetUserAddressList(userId)
	return
}

func (a *ShopUserAddressService) GetDefaultAddressInfo(userId uint) (addressInfo shop.ShopUserAddress, err error) {
	addressInfo, err = shop.GetUserAddressInfo(userId, 0, 1)
	return
}

func (a *ShopUserAddressService) EditUserAddress(userId uint, param *request.EditUserAddressParam) (err error) {
	var userAddress shop.ShopUserAddress
	var defaultAddress shop.ShopUserAddress

	if userAddress, err = shop.GetUserAddressInfo(userId, param.AddressId, 0); err != nil {
		return errors.New("不存在的用户地址")
	}

	//设置默认地址
	if param.DefaultFlag == 1 {
		defaultAddress, _ = shop.GetUserAddressInfo(userId, 0, 1)
		if defaultAddress != (shop.ShopUserAddress{}) {
			defaultAddress.DefaultFlag = 0
			defaultAddress.UpdateTime = time.Now()
			err = shop.SaveUserAddress("edit", &defaultAddress)
			if err != nil {
				return err
			}
		}
	}

	if err = copier.Copy(&userAddress, &param); err != nil {
		return err
	}
	userAddress.UpdateTime = time.Now()
	return shop.SaveUserAddress("edit", &userAddress)
}
