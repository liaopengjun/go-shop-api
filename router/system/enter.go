package system

type RouterGroup struct {
	UserRouter
	BaseRouter
	MenuRouter
	AuthorityRouter
	SysApiRouter
	UploadRoute
	LoginLogRouter
}
