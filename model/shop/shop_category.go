package shop

import (
	"go-shop-api/global"
	"time"
)

type ShopCategory struct {
	CategoryId    int            `json:"categoryId" gorm:"primarykey;AUTO_INCREMENT"`
	CategoryLevel int            `json:"categoryLevel" gorm:"comment:分类等级"`
	ParentId      int            `json:"parentId" gorm:"comment:父类id"`
	CategoryName  string         `json:"categoryName" gorm:"comment:分类名称"`
	CategoryRank  int            `json:"categoryRank" gorm:"comment:排序比重"`
	IsDeleted     int            `json:"isDeleted" gorm:"comment:是否删除"`
	Children      []ShopCategory `json:"children" gorm:"-"`
	CreateTime    time.Time      `json:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"` // 创建时间
	UpdateTime    time.Time      `json:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"` // 更新时间
}

type ShopCategoryData struct {
	CategoryId    int                `json:"categoryId"`
	CategoryLevel int                `json:"categoryLevel"`
	ParentId      int                `json:"parentId"`
	CategoryName  string             `json:"categoryName" `
	Children      []ShopCategoryData `json:"children" gorm:"-"`
}

func (ShopCategory) TableName() string {
	return "shop_categorys"
}

func GetGoodsCategoryList() (category []ShopCategoryData, err error) {
	err = global.GA_DB.Table("shop_categorys").Select("category_id,category_level,parent_id,category_name").Where("is_deleted = 0").Order("category_rank desc").Find(&category).Error
	return
}
