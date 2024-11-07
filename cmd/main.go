package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"shop-grpc-golang/app/shop/controller"
)

func main() {
	//bctx, cancelFunc := context.WithCancel(context.Background())
	//group, gctx := errgroup.WithContext(bctx)
	//donChan := make(chan struct{}, 2)
	//
	//shopRedis := external.NewRedis(bctx)
	//shopMysqlDB := external.NewMysql()

	lis, err := net.Listen("tcp", ":9095")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	controller.RegisterShopHandler(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
