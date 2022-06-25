package system

import (
	"go-admin/global"
	"go-admin/model/system"
	Request "go-admin/model/system/request"
	"go-admin/utils"
	"strconv"
	"strings"
)

type SysApiService struct {
}

func (a *SysApiService) AddSysApi(p *Request.ParamSysApiData) (err error) {
	SysApi := &system.SysApi{
		Path:        p.Path,
		Method:      p.Method,
		ApiGroup:    p.ApiGroup,
		Description: p.Description,
	}
	return system.AddSysApi(SysApi)
}

func (a *SysApiService) UpdateSysApi(p *Request.ParamSysApiData) (err error) {
	//校验是否存在api
	oldSysApi, err := system.ExitApi(p.Id, p.Path, p.Method)
	if err != nil {
		return err
	}
	//更新已经分配的权限
	var casbinService CasbinService
	err = casbinService.UpdateCasbinApi(oldSysApi.Path, p.Path, oldSysApi.Method, p.Method)
	if err != nil {
		return err
	}
	SysApi := &system.SysApi{
		GA_MODEL:    global.GA_MODEL{ID: p.Id},
		Path:        p.Path,
		Method:      p.Method,
		ApiGroup:    p.ApiGroup,
		Description: p.Description,
	}
	return system.UpdateSysApi(SysApi)
}

func (a *SysApiService) DelSysApi(id string) (err error) {
	strids := strings.Split(id, ",")
	var ids []int
	for _, id := range strids {
		if id == "" {
			continue
		}
		idint, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		ids = append(ids, idint)
	}
	return system.DelSysApi(ids)
}

func (a *SysApiService) GetSysApiList(p *Request.ParamSysApiList) (list interface{}, total int64, err error) {
	return system.GetSysApiList(p.Path, p.Method, p.ApiGroup, int(p.Limit), int(p.Page))
}

func (a *SysApiService) GetSysApiAll() (list interface{}, err error) {
	apiAll, err := system.GetSysApiAll()

	//组名
	var keyName []string
	for _, api := range apiAll {
		if len(keyName) == 0 {
			keyName = append(keyName, api.ApiGroup)
		} else {
			if !utils.IsValueInList(api.ApiGroup, keyName) {
				keyName = append(keyName, api.ApiGroup)
			}
		}
	}
	//组名遍历下级
	var apiRes []system.ApiRes
	for _, name := range keyName {
		var sys_api []system.SysApi
		for _, api := range apiAll {
			if name == api.ApiGroup {
				sys_api = append(sys_api, api)
			}
		}
		apiRes = append(apiRes, system.ApiRes{
			ID:          name,
			Description: name + "组",
			Children:    sys_api,
		})
	}
	return apiRes, err
}
