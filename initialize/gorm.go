package initialize

import (
	"fmt"
	"go-admin/config"
	"go-admin/global"
	"go-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

// Gorm gorm初始化连接
func Gorm() *gorm.DB {
	config := new(config.MySQLConfig)
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.GA_CONFIG.MySQLConfig.User,
		global.GA_CONFIG.MySQLConfig.Password,
		global.GA_CONFIG.MySQLConfig.Host,
		global.GA_CONFIG.MySQLConfig.Port,
		global.GA_CONFIG.MySQLConfig.DB,
	)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.MaxIdleConns) // SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(config.MaxOpenConns) // SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetConnMaxLifetime(time.Hour)        // SetConnMaxLifetime 设置了连接可复用的最大时间。
		return db
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysUser{},
		system.SysMenu{},
		system.SysAuthority{},
		system.SysApi{},
		system.SysLoginLog{},
	)
	if err != nil {
		zap.L().Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
}
