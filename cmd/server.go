package main

import (
	"log"
	"real-time-ranking/internal/daos"
	"real-time-ranking/internal/handlers"
	"real-time-ranking/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "real-time-ranking/docs"
)

func main() {
	if err := godotenv.Load("dev.env"); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	redis := daos.NewRedisClient()
	db := daos.NewPostgresDB()

	videoDAO := daos.NewVideoDAO(db, redis)
	rankingService := services.NewRankingService(videoDAO)
	videoHandler := handlers.NewVideoHandler(rankingService)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/videos/:id/score", videoHandler.UpdateScore)
		api.GET("/videos/top", videoHandler.GetTopVideos)
		api.GET("/users/:id/videos/top", videoHandler.GetUserTopVideos)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
