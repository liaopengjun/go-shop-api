package response

import "errors"

var (
	ErrorMenuExit  = errors.New("菜单已存在")
	ErrorChildExit = errors.New("菜单存在下级")
)

type UserResult struct {
	MenuList interface{} `json:"menuList"`
	BtnList  interface{} `json:"btnList"`
}
