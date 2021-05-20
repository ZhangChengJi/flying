package global

import (
	"flying-config/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	LOG        *zap.Logger
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
)
