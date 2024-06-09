package middlewares

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading enviroment variables in the auth middleware")
		}

		JWT_SECRET := os.Getenv("JWT_SECRET")

		authHeader := c.Get("Authorization")
		
		if authHeader == ""{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token missing",
			})
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error)  {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWT_SECRET), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				if time.Unix(int64(exp), 0).Before(time.Now()){
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error": "Token expired",
					})
				}
			}

			c.Locals("userId", claims["id"])

			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}
}