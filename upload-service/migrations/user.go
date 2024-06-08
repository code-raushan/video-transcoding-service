package main

import (
	"log"

	"github.com/code-raushan/video-transcoding-service/upload-service/database"
	"github.com/code-raushan/video-transcoding-service/upload-service/models"
)

func main() {
	database.ConnectDB()

	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
