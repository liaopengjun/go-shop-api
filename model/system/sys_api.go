package system

import (
	"errors"
	"go-shop-api/global"
	"gorm.io/gorm"
)

type SysApi struct {
	global.GA_MODEL
	Path        string `json:"path" gorm:"comment:api路径"`
	Method      string `json:"method" gorm:"default:POST;comment:方法:创建POST(默认)|查看GET|更新PUT|删除DELETE"`
	ApiGroup    string `json:"api_group" gorm:"comment:api组"`
	Description string `json:"description" gorm:"comment:api中文描述"`
}

type ApiRes struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Children    []SysApi `json:"children"`
}

func GetSysApiList(Path, Method, ApiGroup string, Limit, Page int) (sysapi []SysApi, total int64, err error) {
	offset := Limit * (Page - 1)
	db := global.GA_DB.Model(&SysApi{})
	if Path != "" {
		db = db.Where("`path` LIKE ?", "%"+Path+"%")
	}
	if Method != "" {
		db = db.Where("`method` = ?", Method)
	}
	if ApiGroup != "" {
		db = db.Where("`api_group` = ?", ApiGroup)
	}
	err = db.Count(&total).Error
	err = db.Offset(offset).Limit(Limit).Find(&sysapi).Error
	return sysapi, total, err
}

func GetSysApiAll() (sysapi []SysApi, err error) {
	err = global.GA_DB.Find(&sysapi).Error
	return sysapi, err
}

func ExitApi(id uint, newPath string, newMethod string) (oldA *SysApi, err error) {
	err = global.GA_DB.Where("id = ?", id).First(&oldA).Error
	if oldA.Path != newPath || oldA.Method != newMethod {
		if !errors.Is(global.GA_DB.Where("path = ? AND method = ?", newPath, newMethod).First(&SysApi{}).Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("存在相同api路径")
		}
	}
	return oldA, nil
}

func AddSysApi(data *SysApi) (err error) {
	return global.GA_DB.Create(&data).Error
}

func UpdateSysApi(data *SysApi) (err error) {
	return global.GA_DB.Updates(&data).Error
}

func DelSysApi(ids []int) (err error) {
	api := new(SysApi)
	return global.GA_DB.Delete(&api, ids).Error
}
