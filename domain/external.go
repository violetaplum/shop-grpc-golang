package domain

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ExternalMysql interface {
	ShopCore() *gorm.DB
}

type ExternalRedis interface {
	ShopRedis() *redis.Client
}
