package handlers

import (
	"github.com/code-raushan/video-transcoding-service/upload-service/database"
	"github.com/code-raushan/video-transcoding-service/upload-service/models"
	"github.com/code-raushan/video-transcoding-service/upload-service/types"
	"github.com/code-raushan/video-transcoding-service/upload-service/utils"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var req types.SignUpRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to hash the password",
		})
	}

	user := models.User{
		FullName: req.FullName,
		Email: req.Email,
		Password: hashedPassword,
	}

	checkResult := database.DB.First(&user, "email=?", user.Email)

	if checkResult.RowsAffected == 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	dbResult := database.DB.Create(&user)
	if dbResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}