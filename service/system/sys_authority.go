package system

import (
	"go-shop-api/model/system"
	Request "go-shop-api/model/system/request"
	"go-shop-api/model/system/response"
	"time"
)

type AuthorityService struct {
}

// AddAuthority 添加角色
func (a *AuthorityService) AddAuthority(p *Request.ParamAuthorityData) (err error) {
	//1.校验角色是否存在
	if err = system.ExitAuthority(p.AuthorityId, p.AuthorityName); err != nil {
		return err
	}
	//2.创建菜单
	auth := system.SysAuthority{
		AuthorityId:   p.AuthorityId,
		AuthorityName: p.AuthorityName,
		ParentId:      p.ParentId,
		Status:        &p.Status,
	}
	if err = system.CreateAuthority(&auth); err != nil {
		return err
	}
	return
}

// DelAuthority 删除角色
func (a *AuthorityService) DelAuthority(AuthorityId string) (err error) {
	//校验是否还有下级角色
	childAuthList := system.ExitChildAuthority(AuthorityId)
	if len(childAuthList) != 0 {
		return response.ErrorAuthChildExit
	}
	return system.DelAuthority(AuthorityId)
}

// GetAuthorityList 获取角色列表
func (a *AuthorityService) GetAuthorityList(p *Request.ParamAuthorityList) (list interface{}, total int64, err error) {
	allAuthList, total, err := system.GetAuthorityList(p.AuthorityName, p.Status, int(p.Page), int(p.Limit))
	if err != nil {
		return nil, 0, err
	}
	treeMap := make(map[string][]system.SysAuthority)
	for _, v := range allAuthList {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	//递归查询顶级栏目子菜单
	var authList []system.SysAuthority
	authList = treeMap["0"]
	for i := 0; i < len(authList); i++ {
		err = a.getAuthChildrenList(&authList[i], treeMap)
		if err != nil {
			break
		}
	}
	return authList, total, err
}

func (a *AuthorityService) getAuthChildrenList(auth *system.SysAuthority, treeMap map[string][]system.SysAuthority) (err error) {
	auth.Children = treeMap[auth.AuthorityId]
	for i := 0; i < len(auth.Children); i++ {
		err = a.getAuthChildrenList(&auth.Children[i], treeMap)
	}
	return err
}

func (a *AuthorityService) UpdateAuthorityStatus(auth *Request.GetAuthorityId) (err error) {
	return system.UpdateAuthorityStatus(auth.AuthorityId, auth.Status)
}

//GetAuthorityInfo 角色详情
func (a *AuthorityService) GetAuthorityInfo(authority_id string) (auth *system.SysAuthority) {
	return system.GetAuthorityInfo(authority_id)
}

// UpdateAuthority 更新角色
func (a *AuthorityService) UpdateAuthority(p *Request.ParamAuthorityData) (err error) {
	auth := system.SysAuthority{
		UpdatedAt:     time.Time{},
		AuthorityId:   p.AuthorityId,
		AuthorityName: p.AuthorityName,
		ParentId:      p.ParentId,
		Status:        &p.Status,
	}
	return system.UpdateAuthority(&auth)
}

// SetAuthority 设置角色菜单
func (a *AuthorityService) SetAuthority(menus []system.SysMenu, authorityId string) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysMenus = menus
	return system.SetAuthority(&auth)
}
