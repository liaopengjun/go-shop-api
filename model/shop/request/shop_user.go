package request

//  ShopUserParam 用户登陆注册
type ShopUserParam struct {
	UserName string `json:"user_name" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}
