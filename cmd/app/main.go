package main

import (
	"context"
	"crud-grpc-server/internal/handlers"
	"crud-grpc-server/internal/utils"
	"flag"
	"fmt"
	"net"

	gen "crud-grpc-server/third_party/grpcGenerated"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"net/http"
)

var hand *handlers.Handlers

func startGrpc(handlers gen.ServiceExampleServer) {

	port := utils.GetEnv("EXAMPLE_SERVER_GRPC_PORT", "9090").(string)

	server := grpc.NewServer()
	gen.RegisterServiceExampleServer(server, handlers)

	for {
		fmt.Printf("Server grpc started listening... (port:%s)\n", port)

		list, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			println(err)
			continue
		}

		err = server.Serve(list)
		if err != nil {
			println(err)
			continue
		}
	}
}

func startRestAdapter() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	port := utils.GetEnv("EXAMPLE_SERVER_HTTP_PORT", "8081").(string)
	portGRPC := utils.GetEnv("EXAMPLE_SERVER_HTTP_PORT", "9090").(string)
	grpcEndpoint := fmt.Sprintf("localhost:%s", portGRPC)

	grpcServerEndpoint := flag.String("grpc-server-endpoint", grpcEndpoint, "gRPC server endpoint")

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gen.RegisterServiceExampleHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	fmt.Printf("Server http started listening... (port:%s)\n", port)

	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func main() {

	handler, err := handlers.Start()
	if err != nil {
		fmt.Printf("EXIT! Error starting the handler. Err(%v)\n", err)
		return
	}

	go startGrpc(handler)

	if err := startRestAdapter(); err != nil {
		fmt.Printf("EXIT! Error starting the rest adapter. Err(%v)\n", err)
		return
	}
}
