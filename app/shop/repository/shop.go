package repository

import (
	"gorm.io/gorm"
	"shop-grpc-golang/domain"
)

type shopRepository struct {
	storeMysql *gorm.DB
}

func NewShopRepository(db *gorm.DB) domain.ShopRepository {
	return &shopRepository{storeMysql: db}
}
