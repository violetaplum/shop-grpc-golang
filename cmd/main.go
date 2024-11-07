package main

import "shop-grpc-golang/app/shop/controller"

func main() {
	//bctx, cancelFunc := context.WithCancel(context.Background())
	//group, gctx := errgroup.WithContext(bctx)
	//donChan := make(chan struct{}, 2)
	//
	//shopRedis := external.NewRedis(bctx)
	//shopMysqlDB := external.NewMysql()
	controller.RegisterShopHandler()
}
