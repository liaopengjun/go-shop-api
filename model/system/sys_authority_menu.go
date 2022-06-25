package system

type SysAuthorityMenu struct {
	SysMenu
	MenuId      string    `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId string    `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu `json:"children" gorm:"-"`
}

func (s SysAuthorityMenu) TableName() string {
	return "authority_menu"
}
