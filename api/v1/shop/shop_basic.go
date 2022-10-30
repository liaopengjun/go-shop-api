package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ShopBasicApi struct {
}

// IndexInfo 首页
func (b *ShopBasicApi) IndexInfo(g *gin.Context) {
	fmt.Println("首页")
}
