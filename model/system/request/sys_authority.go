package request

import "go-shop-api/model/system"

//列表
type ParamAuthorityList struct {
	Page          int64  `json:"page" binding:"numeric"`
	Limit         int64  `json:"limit" binding:"required,numeric"`
	Sort          string `json:"sort"`
	Status        string `json:"status"`
	AuthorityName string `json:"authority_name"`
}

// 表单参数
type ParamAuthorityData struct {
	AuthorityId   string `json:"authorityId" binding:"required"`
	AuthorityName string `json:"authorityName" binding:"required"`
	ParentId      string `json:"parentId"`
	Status        int    `json:"status"`
}

type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" binding:"required"` // 角色ID
	Status      string `json:"status" form:"status"`           // 角色ID
}

//设置权限
type AddMenuAuthorityInfo struct {
	Menus       []system.SysMenu `json:"menus" binding:"required"`
	AuthorityId string           `json:"authorityId" binding:"required"` // 角色ID
}

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// 设置角色api参数
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId" binding:"required"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}
