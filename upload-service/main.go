package main

import (
	"fmt"
	"log"

	"github.com/code-raushan/video-transcoding-service/upload-service/database"
	"github.com/code-raushan/video-transcoding-service/upload-service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main(){
	app := fiber.New()

	app.Use(logger.New())


	fmt.Println("connecting to database")
	database.ConnectDB()
	fmt.Println("connected to database")

	apiV1 := app.Group("/api/v1")

	routes.HandleV1Routes(&apiV1)
	
	if err := app.Listen(":8500"); err != nil {
		log.Fatalf("Failed to start HTTP server at port 8500")
	}
}

