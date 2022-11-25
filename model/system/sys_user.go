package system

import (
	"github.com/satori/go.uuid"
	"go-shop-api/global"
	"gorm.io/gorm"
)

type SysUser struct {
	global.GA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"comment:用户UUID"`
	Username    string         `json:"userName" gorm:"comment:用户登录名"`
	Password    string         `json:"password"  gorm:"comment:用户登录密码"`
	NickName    string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Sex         string         `json:"sex" gorm:"size:255;comment:性别0:未知1:男2:女"`
	HeaderImg   string         `json:"headerImg" gorm:"default:https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif;comment:用户头像"` // 用户头像
	Email       string         `json:"email" gorm:"size:128;comment:邮箱"`
	Phone       string         `json:"phone" gorm:"size:11;comment:手机号"`
	Salt        string         `json:"-" gorm:"size:255;comment:加盐"`
	Status      string         `json:"status" gorm:"size:4;default:0;comment:状态0:默认正常1:已禁用"`
	AuthorityId string         `json:"authorityId" gorm:"default:1;comment:用户角色ID"`
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

// CheckUserExist 校验用户是否已经注册
func CheckUserExist(username string) (user *SysUser, row int64) {
	row = global.GA_DB.Where("username = ? and status = ?", username, 0).First(&user).RowsAffected
	return
}

func GetUserInfo(uuid uuid.UUID) (user *SysUser, err error) {
	err = global.GA_DB.Preload("Authorities").Preload("Authority").Where("uuid  = ? ", uuid).First(&user).Error
	return user, err
}

// InsertUser 注册用户
func InsertUser(u *SysUser) (err error) {
	return global.GA_DB.Create(&u).Error
}

// UserList 用户列表
func UserList(Page int, Limit int, Status string, Username string, Phone string, AuthorityId string) (user []SysUser, total int64, err error) {
	offset := (Page - 1) * Limit
	db := global.GA_DB.Model(&SysUser{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if Username != "" {
		db = db.Where("`username` LIKE ?", "%"+Username+"%")
	}
	if Phone != "" {
		db = db.Where("`phone` = ?", Phone)
	}
	if Status != "" {
		db = db.Where("`status` = ?", Status)
	}
	if AuthorityId != "" {
		db = db.Where("`authority_id` = ?", AuthorityId)
	}
	var userList []SysUser
	err = db.Count(&total).Error
	err = db.Offset(offset).Limit(Limit).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

func DelUser(ids []int) (err error) {
	user := new(SysUser)
	return global.GA_DB.Delete(&user, ids).Error
}

func UpdateUser(user *SysUser) (err error) {
	return global.GA_DB.Updates(&user).Error
}

func SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.GA_DB.Debug().Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []SysUseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, SysUseAuthority{
				id, v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		var user = SysUser{
			GA_MODEL:    global.GA_MODEL{ID: id},
			AuthorityId: authorityIds[0],
		}
		TxErr = tx.Updates(&user).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

func EditPassword(user *SysUser, NewPassword string) (err error) {
	return global.GA_DB.Where("id = ? and password = ?", user.ID, user.Password).First(&user).Update("password", NewPassword).Error
}
