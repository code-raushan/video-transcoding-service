package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HandleV1Routes(appRouter *fiber.Router) {
	router := *appRouter
	
	router.Route("/auth", UsersRoute())
	router.Route("/upload", UploadRoute())
}