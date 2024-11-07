package domain

import "gorm.io/gorm"

type ExternalMysql interface {
	ShopCore() *gorm.DB
}
