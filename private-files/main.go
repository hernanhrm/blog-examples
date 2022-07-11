package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	loadEnv()

	s3Service, err := NewS3Service(Config{
		S3AccessKey:    os.Getenv("S3_ACCESS_KEY_ID"),
		S3SecretKey:    os.Getenv("S3_SECRET_ACCESS_KEY"),
		S3BucketName:   os.Getenv("S3_BUCKET_NAME"),
		S3BucketRegion: os.Getenv("S3_BUCKET_REGION"),
	})
	if err != nil {
		log.Fatalf("NewS3Service(): %v", err)
	}

	handler := newHandler(NewFileManager(s3Service))

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// these routes should be protected by a middleware that checks
	// the Authentication and Authorization of the user to access the resources

	apiGroup := e.Group("api/v1/filemanager")
	apiGroup.POST("", handler.Upload)

	apiGroup.GET("", handler.GetFile)
	apiGroup.GET("/sign", handler.Presign)

	employeeGroup := e.Group("api/v1/employees")
	employeeGroup.GET("", handler.GetEmployees)

	log.Fatal(e.Start(":8080"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("godotenv.Load(): Error loading .env file: %v", err)
	}
}
