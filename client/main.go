package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fudute/api/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cc, err := grpc.Dial(":8080",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	cli := hello.NewHelloServiceClient(cc)
	resp, err := cli.Hello(context.Background(), &hello.HelloReq{
		Msg: "Hi",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("resp.Msg: %v\n", resp.Msg)
}
