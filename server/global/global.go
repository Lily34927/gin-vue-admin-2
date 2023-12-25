package global

import (
	"server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_Redis  *redis.Client
	GVA_Config config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
)
