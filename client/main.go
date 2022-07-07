package main

import (
	"context"
	"fmt"
	"log"

	hellopb "github.com/fudute/api/pbgen/go/proto/hello/v1"
	"github.com/go-resty/resty/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	byrpc()
}

func byrpc() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	cli := hellopb.NewHelloServiceClient(cc)

	md := metadata.New(map[string]string{
		"from": "grpc",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := cli.GetWeather(ctx, &hellopb.GetWeatherRequest{})
	if err != nil {
		log.Fatal(err)
	}
	switch resp.Weather.(type) {
	case *hellopb.GetWeatherResponse_Raining:
		log.Println("It's rainning day")
	case *hellopb.GetWeatherResponse_Cloudy_:
		log.Println("It's cloudy day")
	}
}

func byhttp() {
	r := resty.New().R()
	resp, err := r.Get("http://localhost:9090/v1/time")
	if err != nil {
		log.Fatal(err)
	}
	body := &hellopb.GetTimeResponse{}
	if err := protojson.Unmarshal(resp.Body(), body); err != nil {
		log.Fatal(err)
	}
	fmt.Println(body.Now.AsTime().Local())
}
