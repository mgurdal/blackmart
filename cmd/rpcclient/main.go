package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mgurdal/blackmarkt/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:4004", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		return
	}

	client := proto.NewUserClient(conn)
	g := gin.Default()
	g.POST("sell/", func(ctx *gin.Context) {
		var payload map[string]string
		if err := ctx.BindJSON(&payload); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "parse heartbeat from http post request error",
			})
			return
		}

		item := &proto.Item{
			Name:     payload["name"],
			Quantity: 1,
		}
		if resp, err := client.MoveToMarket(ctx, item); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"Result": resp.Result})
		}

	})
	if err := g.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("Failed server %v", err)
	}
}
