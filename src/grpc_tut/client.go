package main

import (
	"context"
	"log"

	chat "github.com/kartik-dutt/Learning-Go/chat"
	grpc "google.golang.org/grpc"
)

// Working : build and run server.go
// Run client.go
func main() {
	var conn *grpc.ClientConn
	conn, _ = grpc.Dial(":9000", grpc.WithInsecure())
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	message := chat.Message{
		Body: "Hello from the client.",
	}

	res, _ := c.SayHello(context.Background(), &message)
	log.Printf(res.Body)
}
