package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ShopHomeApi struct {
}

// IndexInfo 首页
func (b *ShopHomeApi) IndexInfo(g *gin.Context) {
	fmt.Println("首页")
}

// ShopInfo 商信息
func (b *ShopHomeApi) ShopInfo(g *gin.Context) {
	fmt.Println("商城信息")
}
