package system

import (
	"go-shop-api/model/system"
	"go-shop-api/model/system/request"
	"go-shop-api/model/system/response"
)

type MenuService struct {
}

//AddMenu 添加菜单
func (m *MenuService) AddMenu(menu *system.SysMenu) (err error) {
	//1.查询菜单是否存在
	if err = system.ExitMenu(menu.Name, 0); err != nil {
		return err
	}
	//2.创建菜单
	if err = system.AddMenu(menu); err != nil {
		return err
	}
	return nil
}

//DelMenu 删除菜单
func (m *MenuService) DelMenu(id int) (err error) {
	//1.检查下级是否未删除
	childMenuList := system.ExitChildMenu(id)
	if len(childMenuList) != 0 {
		return response.ErrorChildExit
	}
	return system.DelMenu(id)
}

//UpdateMenu 更新菜单
func (m *MenuService) UpdateMenu(menu *system.SysMenu) (err error) {
	//检查菜单名称是否存在
	oldMenu := system.GetMenuInfo(int(menu.ID))
	if oldMenu.Name != menu.Name {
		if err = system.ExitMenu(menu.Name, menu.ID); err != nil {
			return err
		}
	}
	if err = system.UpdateMenu(menu); err != nil {
		return err
	}
	return
}

//GetMenuList 获取菜单列表
func (m *MenuService) GetMenuList(p *request.ParamMenuList) (list interface{}, total int64, err error) {
	//获取所有菜单
	allMenuList, total, err := system.GetMenuList(p.Title, int(p.Page), int(p.Limit))
	if err != nil {
		return nil, 0, err
	}

	treeMap := make(map[int][]system.SysMenu)
	for _, v := range allMenuList {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}

	//递归查询顶级栏目子菜单
	var menuList []system.SysMenu
	if len(treeMap[0]) == 0 {
		menuList = allMenuList
	} else {
		menuList = treeMap[0]
		for i := 0; i < len(menuList); i++ {
			err = m.getBaseChildrenList(&menuList[i], treeMap)
			if err != nil {
				break
			}
		}
	}
	return menuList, total, err
}

func (m *MenuService) getBaseChildrenList(menu *system.SysMenu, treeMap map[int][]system.SysMenu) (err error) {
	menu.Children = treeMap[int(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		err = m.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (m *MenuService) GetMenuInfo(id int) (menu *system.SysMenu) {
	//检查菜单名称是否存在
	return system.GetMenuInfo(id)
}

func (m *MenuService) GetMenuTreeList() (list interface{}, total int64, err error) {
	//获取所有菜单
	allMenuList, total, err := system.GetMenuTreeList()
	if err != nil {
		return nil, 0, err
	}
	treeMap := make(map[int][]system.SysMenu)
	for _, v := range allMenuList {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	//递归查询顶级栏目子菜单
	var menuList []system.SysMenu
	menuList = treeMap[0]
	for i := 0; i < len(menuList); i++ {
		err = m.getBaseChildrenList(&menuList[i], treeMap)
		if err != nil {
			break
		}
	}
	return menuList, total, err
}

// GetAuthorityMenuList 设置角色菜单
func (m *MenuService) GetAuthorityMenuList(info *request.GetAuthorityId) (err error, menu []system.SysMenu) {
	return system.GetMenuAuthority(info.AuthorityId)
}

// GetUserMenuList 侧边栏菜单
func (m *MenuService) GetUserMenuList(authorityId string) (menus []system.SysMenu, err error) {
	err, menuTree := system.GetMenuAuthority(authorityId)
	treeMap := make(map[int][]system.SysMenu)
	for _, v := range menuTree {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	menus = treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = m.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}
