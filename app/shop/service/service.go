package service

import "shop-grpc-golang/domain"

type shopService struct {
	shopRepo domain.ShopRepository
}

func NewShopService(r domain.ShopRepository) domain.ShopService {
	return &shopService{shopRepo: r}
}
