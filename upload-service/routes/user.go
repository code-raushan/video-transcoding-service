package routes

import (
	"github.com/code-raushan/video-transcoding-service/upload-service/handlers"
	"github.com/gofiber/fiber/v2"
)

func UsersRoute() func(router fiber.Router ){
	return func(router fiber.Router){
		router.Post("/signup", handlers.SignUp)
		router.Post("/signin", handlers.SignIn)
	}
}