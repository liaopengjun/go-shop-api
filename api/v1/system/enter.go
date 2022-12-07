package system

import "go-shop-api/service"

type ApiGroup struct {
	BaseApi
	MenuApi
	AuthorityApi
	SysApi
	UploadApi
	LoginLogApi
	JobApi
}

var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var authService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var sysApiService = service.ServiceGroupApp.SystemServiceGroup.SysApiService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
var uploadService = service.ServiceGroupApp.SystemServiceGroup.UploadService
var loginLogService = service.ServiceGroupApp.SystemServiceGroup.LoginLogService
var jobService = service.ServiceGroupApp.SystemServiceGroup.JobService
