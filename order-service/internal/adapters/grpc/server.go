package grpc

import (
	"fmt"
	"log"
	"net"

	ports "order-service/internal/ports"
	pb "order-service/proto/gen/go"

	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

type GrpcAdapter struct {
	api  ports.GrpcPort
	port int

	pb.UnimplementedOrderServiceServer
}

func NewGrpcAdapter(api ports.GrpcPort, port int) *GrpcAdapter {
	return &GrpcAdapter{api: api, port: port}
}

func (a GrpcAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, a)
	reflection.Register(grpcServer)
	fmt.Printf("Before starting grpc server !!!")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port, %d ", a.port)
	}
	fmt.Printf("serving grpc on port %d", a.port)
}
