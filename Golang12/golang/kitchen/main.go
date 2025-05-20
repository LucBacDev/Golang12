package main

import (
	"log"
	"source-base-go/golang/proto/transaction"
	"source-base-go/golang/kitchen/api/handler"
	"source-base-go/golang/kitchen/grpcclient"
	"source-base-go/golang/kitchen/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main (){
	app := gin.New()
	conn := grpcclient.NewGRPCClient(":9000")
	defer conn.Close()
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})

	app.Use(crs)
	transactionClient := transaction.NewTransactionServiceClient(conn)
	transactionService := usecase.NewOrderService(transactionClient)
	handler.MakeHandlers(app, transactionService)

	log.Println("Server running on :1000")
	if err := app.Run(":1000"); err != nil {
		log.Fatal(err)
	}

	
}