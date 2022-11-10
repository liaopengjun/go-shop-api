package shop

import (
	"go-admin/global"
	"time"
)

type ShopCategory struct {
	CategoryId    int       `json:"categoryId" gorm:"primarykey;AUTO_INCREMENT"`
	CategoryLevel int       `json:"categoryLevel" gorm:"comment:分类等级"`
	ParentId      int       `json:"parentId" gorm:"comment:父类id"`
	CategoryName  string    `json:"categoryName" gorm:"comment:分类名称"`
	CategoryRank  int       `json:"categoryRank" gorm:"comment:排序比重"`
	IsDeleted     int       `json:"isDeleted" gorm:"comment:是否删除"`
	CreateTime    time.Time `json:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"` // 创建时间
	UpdateTime    time.Time `json:"updateTime" gorm:"column:update_time;comment:修改时间;type:datetime"` // 更新时间
}

func (ShopCategory) TableName() string {
	return "shop_categorys"
}

func GetGoodsCategoryList() (category []*ShopCategory, err error) {
	err = global.GA_DB.Select("category_id", "category_level", "category_name", "parent_id").
		Where("is_deleted = 0").Order("category_rank desc").Find(&category).Error
	return
}
