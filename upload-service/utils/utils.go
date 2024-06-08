package utils

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(passwordBytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateJWT(c *fiber.Ctx, id string) (string, error) {
	if err:= godotenv.Load(); err !=nil {
		log.Fatal("error loading environment variables")
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to sign the token",
		})
	}

	return tokenString, nil
}