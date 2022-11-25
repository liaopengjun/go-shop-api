package service

import (
	"go-shop-api/service/shop"
	"go-shop-api/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	ShopServiceGroup   shop.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
