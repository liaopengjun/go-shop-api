package router

import (
	"go-shop-api/router/shop"
	"go-shop-api/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
	Shop   shop.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
