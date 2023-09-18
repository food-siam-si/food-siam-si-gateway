package main

import (
	"context"
	"log"
	"net"

	"github.com/food-siam-si/food-siam-si-gateway/src/proto"
	"google.golang.org/grpc"
)

type HelloWorldSrv struct {
	proto.UnimplementedHelloServiceServer
}

func (s *HelloWorldSrv) HelloWorld(ctx context.Context, req *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	return &proto.HelloWorldResponse{
		Message: "Hello " + req.Text,
	}, nil
}

func main() {
	app := grpc.NewServer()

	srv := HelloWorldSrv{}

	proto.RegisterHelloServiceServer(app, &srv)

	lis, err := net.Listen("tcp", ":7777")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := app.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
