package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"
	proto "github.com/kartik-dutt/Learning-Go/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:4040", grpc.WithInsecure())
	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, _ := strconv.ParseUint(ctx.Param("num1"), 10, 64)
		b, _ := strconv.ParseUint(ctx.Param("num2"), 10, 64)
		req := &proto.Request{Num1: int64(a), Num2: int64(b)}
		if res, err := client.Add(ctx, req); err == nil {
			ctx.Json(http.StatusOK, gin.H{"result": fmt.Sprintf(res.Ans)})
		}
	})

	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		a, _ := strconv.ParseUint(ctx.Param("num1"), 10, 64)
		b, _ := strconv.ParseUint(ctx.Param("num2"), 10, 64)
		req := &proto.Request{Num1: int64(a), Num2: int64(b)}
		if res, err := client.Multiply(ctx, req); err == nil {
			ctx.Json(http.StatusOK, gin.H{
				"result": fmt.Sprintf(res.Ans),
			})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server", err)
	}
}
