package shop

import (
	"go-admin/model/shop"
	"go-admin/model/shop/request"
	"go-admin/utils"
	"time"
)

type UserService struct {
}

func (u *UserService) Register(p *request.ShopUserParam) (err error) {
	// 检查用户是否已注册
	err = shop.CheckUser(p.UserName)
	if err != nil {
		return err
	}
	// 注册用户
	shopuser := &shop.ShopUser{
		LoginName:     p.UserName,
		PasswordMd5:   utils.MD5V(p.PassWord),
		IntroduceSign: "不经历风雨怎能见彩虹",
		CreateTime:    time.Now(),
	}
	err = shop.CreateShopUser(shopuser)
	if err != nil {
		return err
	}
	return nil
}
