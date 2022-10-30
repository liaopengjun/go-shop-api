package router

import (
	"go-admin/router/shop"
	"go-admin/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Shop   shop.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
