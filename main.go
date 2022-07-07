package main

import (
	"context"
	"log"
	"net"
	"net/http"

	hellopb "github.com/fudute/api/pbgen/go/proto/hello/v1"
	"github.com/fudute/api/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	mu := runtime.NewServeMux()
	// mu := runtime.NewServeMux()
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	ctx := context.TODO()

	hellopb.RegisterHelloServiceHandlerFromEndpoint(ctx, mu, ":8080", opt)

	go http.ListenAndServe(":9090", printHeader(mu))

	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &service.HelloServiceServer{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.Serve(lis))
}

func printHeader(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("recv header: %v", r.Header)
		handler.ServeHTTP(w, r)
	})
}
