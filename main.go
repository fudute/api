package main

import (
	"context"
	"log"
	"net"

	"github.com/fudute/api/proto/hello"
	"google.golang.org/grpc"
)

type HelloServiceServer struct {
	hello.UnimplementedHelloServiceServer
}

func (HelloServiceServer) Hello(ctx context.Context, req *hello.HelloReq) (*hello.HelloResp, error) {
	return &hello.HelloResp{Msg: req.Msg}, nil
}

func main() {
	s := grpc.NewServer()
	s.RegisterService(&hello.HelloService_ServiceDesc, HelloServiceServer{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(s.Serve(lis))
}
