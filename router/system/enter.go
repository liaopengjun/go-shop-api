package system

type RouterGroup struct {
	UserRouter
	BaseRouter
	MenuRouter
	AuthorityRouter
	ApiRouter
	UploadRoute
	LoginLogRouter
	JobRouter
}
