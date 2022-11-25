package system

import (
	"go-shop-api/global"
	"time"
)

type SysLoginLog struct {
	global.GA_MODEL
	Username      string    `json:"username" gorm:"size:128;comment:用户名"`
	Status        string    `json:"status" gorm:"size:4;comment:状态"`
	Ipaddr        string    `json:"ipaddr" gorm:"size:255;comment:ip地址"`
	LoginLocation string    `json:"loginLocation" gorm:"size:255;comment:归属地"`
	Browser       string    `json:"browser" gorm:"size:255;comment:浏览器"`
	Os            string    `json:"os" gorm:"size:255;comment:系统"`
	Platform      string    `json:"platform" gorm:"size:255;comment:固件"`
	LoginTime     time.Time `json:"loginTime" gorm:"comment:登录时间"`
	Remark        string    `json:"remark" gorm:"size:255;comment:备注"`
	Msg           string    `json:"msg" gorm:"size:255;comment:信息"`
}

func CreateLoginLog(l *SysLoginLog) (err error) {
	return global.GA_DB.Create(&l).Error
}

func GetLoginLogList(UserName string, Status string, Ip string, Page int, Limit int) (logs []SysLoginLog, total int64, err error) {
	offset := Limit * (Page - 1)
	db := global.GA_DB.Model(&SysLoginLog{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if UserName != "" {
		db = db.Where("`username` LIKE ?", "%"+UserName+"%")
	}
	if Status != "" {
		db = db.Where("`status` = ?", Status)
	}
	if Ip != "" {
		db = db.Where("`ipaddr` = ?", Ip)
	}
	err = db.Count(&total).Error
	err = db.Offset(offset).Order("id desc").Limit(Limit).Find(&logs).Error
	return logs, total, err
}

func DelLoginLog(ids []int) (err error) {
	return global.GA_DB.Delete(&SysLoginLog{}, ids).Error
}
