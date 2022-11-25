package shop

import (
	"go-shop-api/model/shop"
)

type ShopCategoryService struct {
}

func (s *ShopCategoryService) GetGoodsCategoryList() (categoryList []shop.ShopCategoryData, err error) {
	category, err := shop.GetGoodsCategoryList()
	treeMap := make(map[int][]shop.ShopCategoryData)
	for _, v := range category {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	//递归查询顶级栏目子菜单
	categoryList = treeMap[0]
	for i := 0; i < len(categoryList); i++ {
		err = s.getCategoryChildrenList(&categoryList[i], treeMap)
		if err != nil {
			break
		}
	}
	return categoryList, err
}

func (a *ShopCategoryService) getCategoryChildrenList(cate *shop.ShopCategoryData, treeMap map[int][]shop.ShopCategoryData) (err error) {
	cate.Children = treeMap[cate.CategoryId]
	for i := 0; i < len(cate.Children); i++ {
		err = a.getCategoryChildrenList(&cate.Children[i], treeMap)
	}
	return err
}
