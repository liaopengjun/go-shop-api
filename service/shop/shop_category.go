package shop

import "go-admin/model/shop"

type ShopCategoryService struct {
}

func (s *ShopCategoryService) GetGoodsCategoryList() (category []*shop.ShopCategory, err error) {
	category, err = shop.GetGoodsCategoryList()
	return
}
