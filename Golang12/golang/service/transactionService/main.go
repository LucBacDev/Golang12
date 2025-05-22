package main

import (
	"log"
	"os"
	"source-base-go/golang/proto/auth"
	"source-base-go/golang/proto/wallet"
	"source-base-go/golang/service/transactionService/api/handler"
	"source-base-go/golang/service/transactionService/config"
	"source-base-go/golang/service/transactionService/grpcclient"
	"source-base-go/golang/service/transactionService/infrastructure/repository"
	"source-base-go/golang/service/transactionService/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.SetConfigFile("golang/service/transactionService/config")
	os.Setenv("TZ", "Etc/GMT")
}

func main (){
	app := gin.New()
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})

	app.Use(crs)

	envConfig := getConfig()

	db, err := repository.ConnectDB(envConfig.Sql)
	if err != nil {
		log.Println(err)
		return 
	}
	connAuth := grpcclient.NewGRPCClient(":9000")
	connWallet := grpcclient.NewGRPCClient(":9002")
	defer connAuth.Close()
	defer connWallet.Close()
	
	transactionRepo := repository.NewTransactionRepository(db)

	authGrpcClient := auth.NewAuthServiceClient(connAuth)
	walletGrpcClient := wallet.NewWalletServiceClient(connWallet)

	walletClient := grpcclient.NewWalletClient(walletGrpcClient)
	authClient := grpcclient.NewAuthClient(authGrpcClient)

	transactionService := usecase.NewOrderService(walletClient, authClient, transactionRepo)
	handler.MakeHandlers(app, transactionService)


	go func() {
		if err := NewGRPCServer(":9001"); err != nil {
			log.Print(err)
		}
	}()

	log.Println("Server running on :4000")
	if err := app.Run(":4000"); err != nil {
		log.Print(err)
	}

	
}
func getConfig() config.EnvConfig {
	return config.EnvConfig{
		Host: config.GetString("host.address"),
		Port: config.GetInt("host.port"),
		Sql: config.SqlConfig{
			Timeout:  config.GetInt("database.sql.timeout"),
			DBName:   config.GetString("database.sql.dbname"),
			Username: config.GetString("database.sql.user"),
			Password: config.GetString("database.sql.password"),
			Host:     config.GetString("database.sql.host"),
			Port:     config.GetString("database.sql.port"),
		},
	}
}
