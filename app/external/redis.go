package external

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"shop-grpc-golang/domain"
)

type externalRedis struct {
	storeRedis *redis.Client
}

func (e externalRedis) ShopRedis() *redis.Client {
	return e.storeRedis
}

func NewRedis(ctx context.Context) domain.ExternalRedis {
	r := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})
	// todo: 컨텍스트 갑자기 왜 쓴거지?
	result, err := r.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	logrus.Info(result)
	return &externalRedis{storeRedis: r}
}
