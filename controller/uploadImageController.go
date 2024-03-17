package controller

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

func UploadImage(c echo.Context) error {
	// Read form data
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	defer form.RemoveAll()

	// Get file from form
	file, handler, err := c.Request().FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	// Check file size
	fileSize := handler.Size
	if fileSize > 2000000 { // 2MB
		return c.String(http.StatusBadRequest, "File size exceeds 2MB limit")
	} else if fileSize < 10240 { // 10KB
		return c.String(http.StatusBadRequest, "File size is less than 10KB")
	}

	// Check file extension
	fileExt := filepath.Ext(handler.Filename)
	if fileExt != ".jpg" && fileExt != ".jpeg" {
		return c.String(http.StatusBadRequest, "Only JPG/JPEG files are allowed")
	}

	// Create file in server temp directory
	tempFile, err := os.Create("/tmp/" + handler.Filename)
	if err != nil {
		return err
	}
	defer tempFile.Close()

	// Copy file data to temp file
	if _, err := io.Copy(tempFile, file); err != nil {
		return err
	}

	// Create new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"), // Specify your AWS region
		Credentials: credentials.NewStaticCredentials("AKIAYFGKTK6VHJB4FQPW", "ZXnztIrywJn7ziR6TawS/DyCKXJYUls7F0QX7PGY", ""),
	})
	if err != nil {
		return err
	}

	// Create S3 service client
	svc := s3.New(sess)

	// Open temp file
	fileToUpload, err := os.Open("/tmp/" + handler.Filename)
	if err != nil {
		return err
	}
	defer fileToUpload.Close()

	// Specify S3 bucket name and file path
	bucketName := "sprint-bucket-public-read"
	filePath := "images/" + handler.Filename

	// Upload file to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
		Body:   fileToUpload,
	})
	if err != nil {
		return err
	}

	// Generate URL for uploaded image
	imageURL := "https://" + bucketName + ".s3.amazonaws.com/" + filePath

	// Return JSON response with image URL
	return c.JSON(http.StatusOK, map[string]string{"imageUrl": imageURL})
}
