package system

import (
	"go-shop-api/global"
	Response "go-shop-api/model/system/response"
	"gorm.io/gorm"
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`                                              // 删除时间
	AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`                            // 角色名
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID"`                               // 父角色ID
	Status          *int           `json:"status" gorm:"size:4;comment:状态0:默认正常1:已禁用"`
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	SysMenus        []SysMenu      `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Children        []SysAuthority `json:"children" gorm:"-"`
}

func ExitAuthority(authority_id, authority_name string) (err error) {
	authority := new(SysAuthority)
	if global.GA_DB.Where("authority_id = ? AND authority_name", authority_id, authority_name).Find(&authority).RowsAffected >= 1 {
		return Response.ErrorAuthExit
	}
	return nil
}

func ExitChildAuthority(authority_id string) (authority []SysAuthority) {
	global.GA_DB.Where("parent_id = ?", authority_id).Find(&authority)
	return
}

func GetAuthorityInfo(authority_id string) (auth *SysAuthority) {
	global.GA_DB.Where("authority_id = ?", authority_id).Find(&auth)
	return
}

func CreateAuthority(a *SysAuthority) (err error) {
	return global.GA_DB.Create(&a).Error
}

func DelAuthority(authorityId string) (err error) {
	auth := new(SysAuthority)
	return global.GA_DB.Where("authority_id = ?", authorityId).Delete(&auth).Error
}

func UpdateAuthority(a *SysAuthority) (err error) {
	return global.GA_DB.Updates(&a).Error
}

func GetAuthorityList(AuthorityName string, status string, page int, pagesize int) (authority []SysAuthority, total int64, err error) {
	offset := pagesize * (page - 1)
	db := global.GA_DB.Model(&SysAuthority{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if AuthorityName != "" {
		db = db.Where("`authority_name` LIKE ?", "%"+AuthorityName+"%")
	}
	if status != "" {
		db = db.Where("`status` = ?", status)
	}
	err = db.Count(&total).Error
	err = db.Offset(offset).Limit(pagesize).Find(&authority).Error
	return authority, total, err
}

func SetAuthority(authority *SysAuthority) (err error) {
	var s SysAuthority
	global.GA_DB.Preload("SysMenus").First(&s, "authority_id = ?", authority.AuthorityId)
	err = global.GA_DB.Model(&s).Association("SysMenus").Replace(&authority.SysMenus)
	return err
}

func UpdateAuthorityStatus(authority_id string, status string) (err error) {
	return global.GA_DB.Model(&SysAuthority{}).Where("authority_id = ?", authority_id).Update("status", status).Error
}
