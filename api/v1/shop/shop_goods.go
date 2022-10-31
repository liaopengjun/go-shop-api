package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ShopGoodsApi struct {
}

// SearchGoodsList 搜索
func (goods *ShopGoodsApi) SearchGoodsList(g *gin.Context) {
	fmt.Println("商品搜索")
}

// GoodCategories 商品分类
func (goods *ShopGoodsApi) GoodCategories(g *gin.Context) {
	fmt.Println("商品分类")
}

// GoodDetail 商品详情
func (goods *ShopGoodsApi) GoodDetail(g *gin.Context) {
	fmt.Println("商品详情")
}

// AddShopCart 添加购物车
func (goods *ShopGoodsApi) AddShopCart(g *gin.Context) {
	fmt.Println("添加购物车")
}

// ShopCartList 购物车
func (goods *ShopGoodsApi) ShopCartList(g *gin.Context) {
	fmt.Println("购物车")
}
