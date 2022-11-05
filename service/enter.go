package service

import (
	"go-admin/service/shop"
	"go-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	ShopServiceGroup   shop.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
