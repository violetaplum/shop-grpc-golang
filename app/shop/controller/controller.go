package controller

import (
	"context"
	grpc_shop "github.com/violetaplum/shop-grpc/proto/public_gen/go/shop"
	"google.golang.org/grpc"
)

type handler struct {
}

func (h handler) AddProduct(ctx context.Context, request *grpc_shop.AddProductRequest) (*grpc_shop.AddProductResponse, error) {
	return &grpc_shop.AddProductResponse{}, nil
}

func (h handler) GetProductList(ctx context.Context, request *grpc_shop.GetProductListRequest) (*grpc_shop.GetProductListResponse, error) {
	return &grpc_shop.GetProductListResponse{}, nil
}

func (h handler) GetProduct(ctx context.Context, request *grpc_shop.GetProductRequest) (*grpc_shop.GetProductResponse, error) {
	return &grpc_shop.GetProductResponse{}, nil
}

func (h handler) UpdateProduct(ctx context.Context, request *grpc_shop.UpdateProductRequest) (*grpc_shop.UpdateProductResponse, error) {
	return &grpc_shop.UpdateProductResponse{}, nil
}

func (h handler) DeleteProduct(ctx context.Context, request *grpc_shop.DeleteProductRequest) (*grpc_shop.DeleteProductResponse, error) {
	return &grpc_shop.DeleteProductResponse{}, nil
}

func (h handler) GetStock(ctx context.Context, request *grpc_shop.GetStockRequest) (*grpc_shop.GetStockResponse, error) {
	return &grpc_shop.GetStockResponse{}, nil
}

func (h handler) AddStock(ctx context.Context, request *grpc_shop.AddStockRequest) (*grpc_shop.AddStockResponse, error) {
	return &grpc_shop.AddStockResponse{}, nil
}

func (h handler) DeleteStock(ctx context.Context, request *grpc_shop.DeleteStockRequest) (*grpc_shop.DeleteStockResponse, error) {
	return &grpc_shop.DeleteStockResponse{}, nil
}

func (h handler) GetLowStockList(ctx context.Context, request *grpc_shop.GetLowStockListRequest) (*grpc_shop.GetLowStockListResponse, error) {
	return &grpc_shop.GetLowStockListResponse{}, nil
}

func (h handler) CreateOrder(ctx context.Context, request *grpc_shop.CreateOrderRequest) (*grpc_shop.CreateOrderResponse, error) {
	return &grpc_shop.CreateOrderResponse{}, nil
}

func (h handler) GetOrder(ctx context.Context, request *grpc_shop.GetOrderRequest) (*grpc_shop.GetOrderResponse, error) {
	return &grpc_shop.GetOrderResponse{}, nil
}

func (h handler) GetOrderList(ctx context.Context, request *grpc_shop.GetOrderListRequest) (*grpc_shop.GetOrderListResponse, error) {
	return &grpc_shop.GetOrderListResponse{}, nil
}

func (h handler) UpdateOrderStatus(ctx context.Context, request *grpc_shop.UpdateOrderStatusRequest) (*grpc_shop.UpdateOrderStatusResponse, error) {
	return &grpc_shop.UpdateOrderStatusResponse{}, nil
}

func RegisterShopHandler(srv *grpc.Server) *handler {
	h := &handler{}
	grpc_shop.RegisterShopServiceServer(srv, h)
	return h
}
