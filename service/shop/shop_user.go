package shop

import (
	"errors"
	"go-admin/global"
	"go-admin/model/shop"
	"go-admin/model/shop/request"
	"go-admin/model/shop/response"
	"go-admin/utils"
	"gorm.io/gorm"
	"time"
)

type ShopUserService struct {
}

func (u *ShopUserService) Register(p *request.ShopUserParam) (err error) {
	// 检查用户是否已注册
	if !errors.Is(global.GA_DB.Where("login_name = ?", p.UserName).First(&shop.ShopUser{}).Error, gorm.ErrRecordNotFound) {
		return response.ErrorUserExit
	}
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

func (s *ShopUserService) Login(p *request.ShopUserParam) (user *shop.ShopUser, err error) {
	// 检查用户
	err = global.GA_DB.Where("login_name = ? ", p.UserName).First(&user).Error
	if err != nil {
		return nil, err
	}

	//密码是否一致
	oldPassword := user.PasswordMd5
	newPassword := utils.MD5V(p.PassWord)
	if oldPassword != newPassword {
		return nil, response.ErrorPasswordWrong
	}
	return user, err
}

func (u *ShopUserService) GetUserInfo(user_id uint) (user *shop.ShopUser, err error) {
	user, err = shop.GetUserDetail(user_id)
	if err != nil {
		return nil, err
	}
	return
}

func (u *ShopUserService) EditUser(p *request.ShopEditUserParam) error {

	// 检查用户
	user := new(shop.ShopUser)
	err := global.GA_DB.Where("login_name = ? ", p.UserName).First(&user).Error
	if err != nil {
		return err
	}

	//oldPassword := user.PasswordMd5
	newPassword := utils.MD5V(p.PassWord)
	//if oldPassword != newPassword {
	//	return response.ErrorPasswordWrong
	//}

	userdata := &shop.ShopUser{
		UserId:        user.UserId,
		LoginName:     p.UserName,
		PasswordMd5:   newPassword,
		IntroduceSign: p.IntroduceSign,
	}
	return shop.UpdateUser(userdata)

}
