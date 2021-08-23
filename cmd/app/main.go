package main

import (
	"context"
	"crud-grpc-server/internal/handlers"
	"flag"
	"github.com/golang/glog"
	"net"

	gen "crud-grpc-server/third_party/grpcGenerate"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"net/http"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func startGrpc() {
	println("start grpc")

	server := grpc.NewServer()
	//handlers := &handlers.Handlers{}
	gen.RegisterServiceExampleServiceServer(server, &handlers.Handlers{})

	for {
		println("grpc listen")
		list, err := net.Listen("tcp", ":9090")
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

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gen.RegisterServiceExampleServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	println("start")
	flag.Parse()
	defer glog.Flush()

	go startGrpc()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
