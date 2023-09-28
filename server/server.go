package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "grpc_yt/gen/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	go func() {
		// Create a new ServeMux for the HTTP server
		mux := runtime.NewServeMux()

		// Register gRPC handler on the ServeMux
		ctx := context.Background()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := pb.RegisterTestApiHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
		if err != nil {
			log.Fatalf("Failed to register gRPC gateway: %v", err)
		}

		// Start the HTTP server
		log.Fatalln(http.ListenAndServe(":8081", mux))
	}()

	// Start the gRPC server
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
