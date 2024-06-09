package utils

import (
	"log"
	"os"
	"time"

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

func CreateJWT(id string) (string, error) {
	if err:= godotenv.Load(); err !=nil {
		log.Fatal("error loading environment variables")
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, _ := token.SignedString([]byte(JWT_SECRET))
	
	return tokenString, nil
}