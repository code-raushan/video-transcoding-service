package routes

import (
	"github.com/code-raushan/video-transcoding-service/upload-service/handlers"
	"github.com/code-raushan/video-transcoding-service/upload-service/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UploadRoute() func(router fiber.Router){
	return func(router fiber.Router){
		router.Post("/presign_url", middlewares.AuthMiddleware(), handlers.CreatePresignedUrl)
	}
}