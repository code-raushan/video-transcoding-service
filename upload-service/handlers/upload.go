package handlers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/code-raushan/video-transcoding-service/upload-service/types"
	"github.com/code-raushan/video-transcoding-service/upload-service/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func CreatePresignedUrl(c *fiber.Ctx) error {
	
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading environment variables")
	}

	userId := c.Locals("userId")

	var req types.UploadParams

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse the JSON",
		})
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error uploading the video",
		})
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)

	file_extn := strings.Split(req.Filename, ".")[1]

	file_suffix, _ := utils.GenerateShortUUID()

	key := fmt.Sprintf("uploads/%s/raw/%s_%s.%s", userId, req.Filename, file_suffix, file_extn)

	presignParams := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_UPLOAD_BUCKET")),
		Key: aws.String(key),
	}

	presign_url, err := presignClient.PresignPutObject(c.Context(), presignParams, s3.WithPresignExpires(15 * time.Minute))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "error creating the presigned url",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Use this presigned url to upload the video",
		"url": presign_url.URL,
	})
}