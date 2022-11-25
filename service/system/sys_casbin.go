package system

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go-shop-api/global"
	"go-shop-api/model/system/request"
	"go-shop-api/model/system/response"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
)

type CasbinService struct {
}

// Casbin 持久化到数据库  引入自定义规则
func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	// 通过现有的 Gorm 实例创建适配器
	a, _ := gormadapter.NewAdapterByDB(global.GA_DB)
	// 通过文件或数据库创建一个同步的执行器
	syncedEnforcer, err := casbin.NewSyncedEnforcer(global.GA_CONFIG.CasbinConfig.ModelPath, a)
	if err != nil {
		panic(err)
	}
	// 从数据库重新加载策略
	err = syncedEnforcer.LoadPolicy()
	if err != nil {
		panic(err)
	}
	return syncedEnforcer
}

// GetPolicyPathByAuthorityId 获取角色权限列表
func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// UpdateCasbin 更新权限
func (casbinService *CasbinService) UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) (err error) {
	// 先清除该角色策略
	casbinService.ClearCasbin(0, authorityId)
	if len(casbinInfos) == 0 {
		return nil
	}
	// 重新添加策略
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return response.ErrorAuthApiExit
	}
	return nil
}

// ClearCasbin 清除权限
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// UpdateCasbinApi更新api
func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GA_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}
