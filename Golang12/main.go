package main

import (
	"fmt"
	"log"
	"os"
	authHandler "source-base-go/api/handler/auth"
	walletHandler "source-base-go/api/handler/wallet"

	"source-base-go/config"
	"source-base-go/infrastructure/repository"
	"source-base-go/infrastructure/repository/util"

	"source-base-go/usecase/user"
	"source-base-go/usecase/wallet"


	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	////////////////GRPC
	// "source-base-go/handler"
    // "source-base-go/grpcclient"

    // "google.golang.org/grpc"

)

func init() {
	config.SetConfigFile("config")
	os.Setenv("TZ", "Etc/GMT")
}

func main() {
	envConfig := getConfig()

	//Database Connect
	db, err := repository.ConnectDB(envConfig.Sql)
	if err != nil {
		log.Println(err)
		return
	}
	app := gin.New()
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})

	app.Use(crs)
	//Handler
	userRepo := repository.NewUserRepository(db) 
	walletRepo := repository.NewWalletRepository(db)

	var verifier util.Verifier = util.NewAccessTokenRepository(db)

	userService := user.NewService(userRepo)
	walletService := wallet.NewService(walletRepo, verifier)


	authHandler.MakeHandlers(app, userService)
	walletHandler.MakeHandlers(app, walletService,verifier)


	addr := fmt.Sprintf("%s:%d", envConfig.Host, envConfig.Port)
	log.Printf(" Server đang chạy tại http://%s", addr)
	for _, route := range app.Routes() {
		log.Printf("API: %s %s", route.Method, route.Path)
	}
	if err := app.Run(addr); err != nil {
		log.Fatalf(" Lỗi khi chạy server: %v", err)
	}



	/////////////////////////////// gRPC
	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    // if err != nil {
    //     log.Fatalf("failed to connect to gRPC: %v", err)
    // }
    // defer conn.Close()

    // authClient := grpcclient.NewAuthGRPCClient(conn)
    // authHandler := handler.NewAuthHandler(authClient)

    // r := gin.Default()
    // r.POST("/auth/verify", authHandler.VerifyToken)
    // r.Run(":8080")
	
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
