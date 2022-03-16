package main

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sonkt1210/tiki-vas/handlers"
	"log"
	"os"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.GET("/ipay/check-balance", handlers.CheckBalance)
		client.POST("/ipay/direct-topup", handlers.DirectTopup)
		client.POST("/ipay/buy-card", handlers.BuyCard)
		client.POST("/ipay/check-transaction", handlers.CheckTransaction)
		client.POST("/ipay/retrieve-card-info", handlers.RetrieveCardInfo)
		client.POST("/ipay/check-product-info", handlers.CheckProductInfo)
	}

	return r
}

func main() {
	r := setupRouter()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	port := os.Getenv("PORT")

	r.Run(":" + port) // Ứng dụng chạy tại cổng 8089
}
