package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go-admin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	GA_VP     *viper.Viper
	GA_LOG    *zap.Logger
	GA_CONFIG config.AppConfig
	GA_DB     *gorm.DB
	GA_REDIS  *redis.Client
	GA_TRANS  ut.Translator
)

type GA_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
