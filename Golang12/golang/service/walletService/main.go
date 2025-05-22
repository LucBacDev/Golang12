package main

import (
	"os"
	"source-base-go/golang/service/walletService/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config.SetConfigFile("golang/service/walletService/config")
	os.Setenv("TZ", "Etc/GMT")
}
func main() {
	app := gin.New()
	crs := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie", "Authorization"},
	})

	app.Use(crs)

	grpcServer := NewGRPCServer(":9002")
	grpcServer.Run()

	
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
