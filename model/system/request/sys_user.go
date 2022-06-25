package request

type Register struct {
	Username     string   `json:"userName" binding:"required"`
	Password     string   `json:"passWord" binding:"required"`
	Sex          string   `json:"sex" binding:"required,oneof=0 1 2"`
	Email        string   `json:"email" binding:"required"`
	Phone        string   `json:"phone" binding:"required,numeric"`
	NickName     string   `json:"nickName" gorm:"default:'junge666'"`
	HeaderImg    string   `json:"headerImg"`
	AuthorityId  string   `json:"authorityId" gorm:"default:1"`
	AuthorityIds []string `json:"authorityIds"`
}

type Login struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

type GetUserList struct {
	Username    string `json:"username"`                         // 用户名
	AuthorityId string `json:"authority_id"`                     //角色
	Phone       string `json:"phone"`                            //手机号
	Status      string `json:"status"`                           //状态
	Page        int64  `json:"page" binding:"numeric"`           //页码
	Limit       int64  `json:"limit" binding:"required,numeric"` //分页数
}

type EditUserParam struct {
	ID           int      `json:"id" binding:"required"`
	Username     string   `json:"userName" binding:"required"`
	Sex          string   `json:"sex" binding:"required,oneof=0 1 2"`
	Email        string   `json:"email" binding:"required"`
	Phone        string   `json:"phone" binding:"required,numeric"`
	HeaderImg    string   `json:"headerImg"`
	NickName     string   `json:"nickName" gorm:"default:'junge666'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:1"`
	AuthorityIds []string `json:"authorityIds"`
	Status       string   `json:"status" gorm:"default:1"`
}

type ChangePasswordStruct struct {
	UserID      int    `json:"user_id" binding:"required"`     // 用户名
	IsOldPwd    string `json:"is_old_pwd"`                     // 用户名
	Password    string `json:"password" binding:"required"`    // 用户名
	NewPassword string `json:"newPassword" binding:"required"` // 新密码
}

type EditUserStatus struct {
	UserId int    `json:"userid" binding:"required"` // 用户名
	Status string `json:"status" binding:"required"` // 密码
}

type DelUserAvaterParam struct {
	UserId    int    `json:"userid" binding:"required"`    // 用户名
	HeaderImg string `json:"headerImg" binding:"required"` // 头像路径
}
