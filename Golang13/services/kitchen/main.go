package main

import (
	"log"
	"source-base-go/common/genproto/orders"
	"source-base-go/services/kitchen/api/handler"
	"source-base-go/services/kitchen/grpc_client"
	"source-base-go/services/kitchen/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func NewGRPCClient(addr string) *grpc.ClientConn{
// 	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 		if err != nil {
// 			log.Fatalf("did not corret: %v", err)
// 		}
// 		return conn
// }
// func main (){
// 	httpServer := NewHttpServer(":1000")
// 	httpServer.Run()
// }

func main (){
	app := gin.New()
	conn := grpc_client.NewGRPCClient(":9000")
	defer conn.Close()
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})

	app.Use(crs)
	orderClient := orders.NewOrderServiceClient(conn)
	orderService := usecase.NewOrderService(orderClient)
	handler.MakeHandlers(app, orderService)

	log.Println("Server running on :1000")
	if err := app.Run(":1000"); err != nil {
		log.Fatal(err)
	}

	
}