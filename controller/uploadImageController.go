package controller

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	uuid "github.com/nu7hatch/gouuid"
)

// func UploadImageController(c echo.Context) error {
// 	var awsAccesKeyId = os.Getenv("S3_ID")
// 	var awsSecretAccessKey = os.Getenv("S3_SECRET_KEY")
// 	var awsBucketName = os.Getenv("S3_BUCKET_NAME")
// 	var awsRegion = "ap-southeast-1"

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return err
// 	}

// 	ext := filepath.Ext(file.Filename)
// 	if ext != ".jpg" && ext != ".jpeg" {
// 		return util.ErrorHandler(c, http.StatusBadRequest, FileFormatNotSupported)
// 	}

// 	if file.Size > 2<<20 {
// 		return util.ErrorHandler(c, http.StatusBadRequest, FileSizeExceedsMaximumAllowedSize)
// 	}
// 	if file.Size < 10<<10 {
// 		return util.ErrorHandler(c, http.StatusBadRequest, FileSizeLessThanMinimumAllowedSize)
// 	}

// 	uuidValue, _ := uuid.NewV4()
// 	filename := uuidValue.String() + filepath.Ext(file.Filename)

// 	fileContent, _ := file.Open()
// 	defer fileContent.Close()

// 	cfg, _ := config.LoadDefaultConfig(context.TODO(),
// 		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccesKeyId, awsSecretAccessKey, "")),
// 		config.WithRegion(awsRegion),
// 	)

// 	client := s3.NewFromConfig(cfg)

// 	uploader := manager.NewUploader(client)

// 	uploadResult, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
// 		Bucket: &awsBucketName,
// 		Key:    &filename,
// 		Body:   fileContent,
// 	})

// 	if err != nil {
// 		return util.ErrorHandler(c, http.StatusInternalServerError, FailedToUploadImage)
// 	}

// 	return util.UploadImageResponseHandler(c, http.StatusOK, uploadResult.Location)
// }

func UploadImage(c echo.Context) error {

	var awsAccesKeyId = os.Getenv("S3_ID")
	var awsSecretAccessKey = os.Getenv("S3_SECRET_KEY")
	var awsBucketName = os.Getenv("S3_BUCKET_NAME")
	var awsRegion = "ap-southeast-1"

	// Read form data
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	return err
	// }
	// defer form.RemoveAll()

	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// Check file extension
	fileExt := filepath.Ext(file.Filename)
	if fileExt != ".jpg" && fileExt != ".jpeg" {
		return c.String(http.StatusBadRequest, "Only JPG/JPEG files are allowed")
	}

	// Check file size
	fileSize := file.Size
	if fileSize > 2000000 { // 2MB
		return c.String(http.StatusBadRequest, "File size exceeds 2MB limit")
	} else if fileSize < 10240 { // 10KB
		return c.String(http.StatusBadRequest, "File size is less than 10KB")
	}

	uuidValue, _ := uuid.NewV4()
	filename := uuidValue.String() + filepath.Ext(file.Filename)

	fileContent, _ := file.Open()
	defer fileContent.Close()

	cfg, _ := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccesKeyId, awsSecretAccessKey, "")),
		config.WithRegion(awsRegion),
	)

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)

	uploadResult, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &awsBucketName,
		Key:    &filename,
		Body:   fileContent,
	})
	if err != nil {
		return err
	}

	// Return JSON response with image URL
	return c.JSON(http.StatusOK, map[string]string{"imageUrl": uploadResult.Location})
}

// func UploadImage(c echo.Context) error {
// 	// Read form data
// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		return err
// 	}
// 	defer form.RemoveAll()

// 	// Get file from form
// 	file, handler, err := c.Request().FormFile("file")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Check file size
// 	fileSize := handler.Size
// 	if fileSize > 2000000 { // 2MB
// 		return c.String(http.StatusBadRequest, "File size exceeds 2MB limit")
// 	} else if fileSize < 10240 { // 10KB
// 		return c.String(http.StatusBadRequest, "File size is less than 10KB")
// 	}

// 	// Check file extension
// 	fileExt := filepath.Ext(handler.Filename)
// 	if fileExt != ".jpg" && fileExt != ".jpeg" {
// 		return c.String(http.StatusBadRequest, "Only JPG/JPEG files are allowed")
// 	}

// 	// Create file in server temp directory
// 	tempFile, err := os.Create("/tmp/" + handler.Filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer tempFile.Close()

// 	// Copy file data to temp file
// 	if _, err := io.Copy(tempFile, file); err != nil {
// 		return err
// 	}

// 	// Create new AWS session
// 	sess, err := session.NewSession(&aws.Config{
// 		Region:      aws.String("ap-southeast-1"), // Specify your AWS region
// 		Credentials: credentials.NewStaticCredentials("AKIAYFGKTK6VHJB4FQPW", "ZXnztIrywJn7ziR6TawS/DyCKXJYUls7F0QX7PGY", ""),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Create S3 service client
// 	svc := s3.New(sess)

// 	// Open temp file
// 	fileToUpload, err := os.Open("/tmp/" + handler.Filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer fileToUpload.Close()

// 	// Specify S3 bucket name and file path
// 	bucketName := "sprint-bucket-public-read"
// 	filePath := "images/" + handler.Filename

// 	// Upload file to S3
// 	_, err = svc.PutObject(&s3.PutObjectInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(filePath),
// 		Body:   fileToUpload,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Generate URL for uploaded image
// 	imageURL := "https://" + bucketName + ".s3.amazonaws.com/" + filePath

// 	// Return JSON response with image URL
// 	return c.JSON(http.StatusOK, map[string]string{"imageUrl": imageURL})
// }
