package system

import (
	"go-shop-api/global"
	"go-shop-api/model/system/response"
)

type SysMenu struct {
	global.GA_MODEL
	Title         string         `json:"title" binding:"required" gorm:"comment:菜单名"`          // 菜单名
	Icon          string         `json:"icon"  binding:"required" gorm:"comment:菜单图标"`         // 菜单图标
	Path          string         `json:"path"  binding:"required" gorm:"comment:路由path"`       // 路由path
	Name          string         `json:"name" binding:"required" gorm:"comment:路由name"`        // 路由name
	ParentId      int            `json:"parentId"  gorm:"comment:父菜单ID"`                       // 父菜单ID
	Hidden        *int           `json:"hidden" binding:"required" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component     string         `json:"component" binding:"required" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort          *int           `json:"sort" binding:"required,numeric" gorm:"comment:排序标记"`  // 排序标记
	SysAuthoritys []SysAuthority `json:"authority"  gorm:"many2many:sys_authority_menus;"`
	Children      []SysMenu      `json:"children" gorm:"-"`
}

func ExitMenu(name string, id uint) (err error) {
	menu := new(SysMenu)
	if global.GA_DB.Where("name = ? and id !=? ", name, id).Find(&menu).RowsAffected >= 1 {
		return response.ErrorMenuExit
	}
	return nil
}

func AddMenu(menu *SysMenu) (err error) {
	return global.GA_DB.Create(&menu).Error
}

func DelMenu(id int) (err error) {
	menu := new(SysMenu)
	return global.GA_DB.Where("id = ?", id).Delete(&menu).Error
}

func GetMenuInfo(id int) (menu *SysMenu) {
	global.GA_DB.Where("id = ?", id).Find(&menu)
	return
}

func ExitChildMenu(pid int) (menu []*SysMenu) {
	global.GA_DB.Where("parent_id = ? ", pid).Find(&menu)
	return
}

func UpdateMenu(menu *SysMenu) (err error) {
	return global.GA_DB.Updates(&menu).Error
}

func GetMenuList(title string, page, pagesize int) (menu []SysMenu, count int64, err error) {
	offset := (page - 1) * pagesize
	db := global.GA_DB.Model(&SysMenu{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if title != "" {
		db = db.Where("`title` LIKE ?", "%"+title+"%")
	}
	err = db.Count(&count).Error
	err = db.Offset(offset).Limit(pagesize).Order("sort desc").Find(&menu).Error
	return menu, count, err
}

func GetMenuTreeList() (menu []SysMenu, count int64, err error) {
	db := global.GA_DB.Model(&SysMenu{})
	err = db.Count(&count).Error
	err = db.Find(&menu).Error
	return menu, count, err
}

func GetMenuAuthority(authorityId string) (err error, menus []SysMenu) {
	auth := SysAuthority{
		AuthorityId: authorityId,
	}
	err = global.GA_DB.Model(&auth).Order("sort desc").Association("SysMenus").Find(&menus)
	return err, menus
}
