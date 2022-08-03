package dd

import (
	"animeii.tech/model"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Cf    model.Config
	DB    *gorm.DB
	Log   *zap.Logger
	Redis *redis.Client
)

func Load() {
	InitConfig()
	InitLog()
	InitGorm()
	InitRedis()
	InitValidator()
}
