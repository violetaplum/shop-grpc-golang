package controller

import (
	"context"
	grpc_shop "github.com/violetaplum/shop-grpc/proto/public_gen/go/shop"
	"google.golang.org/grpc"
)

type handler struct {
	//grpc_shop.UnimplementedShopServiceServer
}

func (h handler) AddProduct(ctx context.Context, request *grpc_shop.AddProductRequest) (*grpc_shop.AddProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetProductList(ctx context.Context, request *grpc_shop.GetProductListRequest) (*grpc_shop.GetProductListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetProduct(ctx context.Context, request *grpc_shop.GetProductRequest) (*grpc_shop.GetProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) UpdateProduct(ctx context.Context, request *grpc_shop.UpdateProductRequest) (*grpc_shop.UpdateProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) DeleteProduct(ctx context.Context, request *grpc_shop.DeleteProductRequest) (*grpc_shop.DeleteProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetStock(ctx context.Context, request *grpc_shop.GetStockRequest) (*grpc_shop.GetStockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) AddStock(ctx context.Context, request *grpc_shop.AddStockRequest) (*grpc_shop.AddStockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) DeleteStock(ctx context.Context, request *grpc_shop.DeleteStockRequest) (*grpc_shop.DeleteStockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetLowStockList(ctx context.Context, request *grpc_shop.GetLowStockListRequest) (*grpc_shop.GetLowStockListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) CreateOrder(ctx context.Context, request *grpc_shop.CreateOrderRequest) (*grpc_shop.CreateOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetOrder(ctx context.Context, request *grpc_shop.GetOrderRequest) (*grpc_shop.GetOrderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetOrderList(ctx context.Context, request *grpc_shop.GetOrderListRequest) (*grpc_shop.GetOrderListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h handler) UpdateOrderStatus(ctx context.Context, request *grpc_shop.UpdateOrderStatusRequest) (*grpc_shop.UpdateOrderStatusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func RegisterShopHandler(srv *grpc.Server) *handler {
	h := &handler{}
	grpc_shop.RegisterShopServiceServer(srv, h)
	return h
}
