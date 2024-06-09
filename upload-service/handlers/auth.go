package handlers

import (
	"log"

	"github.com/code-raushan/video-transcoding-service/upload-service/database"
	"github.com/code-raushan/video-transcoding-service/upload-service/models"
	"github.com/code-raushan/video-transcoding-service/upload-service/types"
	"github.com/code-raushan/video-transcoding-service/upload-service/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
		Email:    req.Email,
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

func SignIn(c *fiber.Ctx) error {
    var req types.SignInRequest

    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse the JSON body",
        })
    }

    var user models.User

    err := database.DB.Where("email = ?", req.Email).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid email or password",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    if !utils.CheckPassword(req.Password, user.Password) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    jwtToken, err := utils.CreateJWT(user.ID)
    if err != nil {
        log.Fatal(err.Error())
    }

    return c.JSON(fiber.Map{
        "message": "Successfully signed in",
        "user":    user,
        "token":   jwtToken,
    })
}

