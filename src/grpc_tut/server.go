package main

import (
	"log"
	"net"

	chat "github.com/kartik-dutt/Learning-Go/chat"
	grpc "google.golang.org/grpc"
)

func main() {
	lis, _ := net.Listen("tcp", ":9000")
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve on port 9000")
	}

}
