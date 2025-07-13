package routes

import (
	"github.com/gofiber/fiber/v2"
)

func HandleV1Routes(appRouter *fiber.Router) {
	router := *appRouter
	
	router.Route("/", func(router fiber.Router) {
		router.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "V1 APIs",
			})
		})
	})
	router.Route("/auth", UsersRoute())
	router.Route("/upload", UploadRoute())
}