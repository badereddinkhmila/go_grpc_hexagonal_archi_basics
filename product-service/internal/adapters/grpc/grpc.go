package grpc

import (
	"context"
	"log"

	pb "product-service/proto/gen/go"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

var (
	// ErrNotFound error type
	ErrNotFound = status.Errorf(codes.NotFound, "Service not found")
	// ErrServiceUnimplemented error type
	ErrServiceUnimplemented = status.Errorf(codes.Unimplemented, "Service not implemented")
)

func (a GrpcAdapter) GetCategories(ctx context.Context, request *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	log.Printf("Payment request for order with id: %d", request.GetOrderId())
	return &pb.PlaceOrderResponse{
		OrderId:   request.GetOrderId(),
		PaymentId: 12345,
	}, nil
}

func (a GrpcAdapter) GetProducts(ctx context.Context, request *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	return &pb.GetOrdersResponse{
		Orders: []*pb.Order{{
			Id:   33,
			Name: "Order 1",
		}, {Id: 44, Name: "Order 2"}},
	}, nil
}
