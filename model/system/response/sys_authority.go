package response

import "errors"

var (
	ErrorAuthExit      = errors.New("角色已存在")
	ErrorAuthChildExit = errors.New("角色权限存在下级")
	ErrorAuthApiExit   = errors.New("存在相同api,添加失败,请联系管理员")
)
