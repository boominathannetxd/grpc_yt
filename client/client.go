package main

import (
	"context"
	pb "grpc_yt/gen/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewTestApiClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{
		Msg: "Hello world",
	})
	if err != nil {
		panic(err)
	}
	println(resp.Msg)

}
