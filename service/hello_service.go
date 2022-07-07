package service

import (
	"context"
	"log"

	hellopb "github.com/fudute/api/pbgen/go/proto/hello/v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HelloServiceServer struct {
	hellopb.UnimplementedHelloServiceServer
}

func (*HelloServiceServer) GetTime(ctx context.Context, req *hellopb.GetTimeRequest) (*hellopb.GetTimeResponse, error) {
	resp := new(hellopb.GetTimeResponse)
	resp.Now = timestamppb.Now()
	return resp, nil
}

func (*HelloServiceServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("get md failed")
	}
	for k, v := range md {
		log.Printf("%v: %v", k, v)
	}
	return &hellopb.HelloResponse{
		HelloMsg: "hello " + req.User.FirstName,
	}, nil
}

func (*HelloServiceServer) Hi(ctx context.Context, req *hellopb.HiRequest) (*hellopb.HiResponse, error) {
	resp := new(hellopb.HiResponse)
	if req.User != nil {
		resp.HiMsg = "hi " + req.User.FirstName
	}
	return resp, nil
}

func (*HelloServiceServer) GetWeather(ctx context.Context, req *hellopb.GetWeatherRequest) (*hellopb.GetWeatherResponse, error) {
	return &hellopb.GetWeatherResponse{
		Weather: &hellopb.GetWeatherResponse_Cloudy_{},
	}, nil
}