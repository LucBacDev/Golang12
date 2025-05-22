package main

import (
	"fmt"
	"log"
	"os"
	"source-base-go/golang/service/userService/config"
	authHandler "source-base-go/golang/service/userService/api/handler/auth"

	"source-base-go/golang/service/userService/infrastructure/repository"

	"source-base-go/golang/service/userService/usecase/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.SetConfigFile("golang/service/walletService/config")
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

	userService := user.NewService(userRepo)

	authHandler.MakeHandlers(app, userService)

	addr := fmt.Sprintf("%s:%d", envConfig.Host, envConfig.Port)
	log.Printf(" Server đang chạy tại http://%s", addr)
	for _, route := range app.Routes() {
		log.Printf("API: %s %s", route.Method, route.Path)
	}
	if err := app.Run(addr); err != nil {
		log.Fatalf(" Lỗi khi chạy server: %v", err)
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
