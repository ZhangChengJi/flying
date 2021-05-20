package global

import (
	"flying-admin/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	LOG        *zap.Logger
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	JWT        config.JWT
	GVA_REDIS  *redis.Client
)
