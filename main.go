package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello")

	// khởi tạo một đối tượng gRPC service
	grpcServer := grpc.NewServer()

	// đăng ký service với grpcServer (của gRPC plugin)
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	// cung cấp gRPC service trên port `1234`
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
