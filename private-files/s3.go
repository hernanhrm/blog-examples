package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Permissions
const (
	PublicRead        = "public-read"
	AuthenticatedRead = "authenticated-read"
	Private           = "private"
)

// Service is the use case of the storage s3 service
type Service struct {
	Bucket  string
	Session *session.Session
}

// NewS3Service returns a new s3 service
func NewS3Service(conf Config) (Service, error) {
	awsSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(conf.S3BucketRegion),
		Credentials: credentials.NewStaticCredentials(conf.S3AccessKey, conf.S3SecretKey, ""),
	}))

	return Service{
		Bucket:  conf.S3BucketName,
		Session: awsSession,
	}, nil
}

// Upload take an io.Reader and uploads it to aws s3
func (s Service) Upload(fileBytes []byte, contentType, path string, isPublic bool) error {
	permission := PublicRead
	if !isPublic {
		permission = Private
	}

	service := s3.New(s.Session)
	_, err := service.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(path),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
		ACL:         aws.String(permission),
	})
	if err != nil {
		return fmt.Errorf("s3.Upload(): %v", err)
	}

	return nil
}

// GetFile returns a file bytes
func (s Service) GetFile(filename string) (GetFileResponse, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	}

	service := s3.New(s.Session)
	result, err := service.GetObject(input)
	if err != nil {
		if aswErr, ok := err.(awserr.Error); ok {
			switch aswErr.Code() {
			case s3.ErrCodeNoSuchKey:
				return GetFileResponse{}, fmt.Errorf("s3.GetFile.service.GetFile(): file does not exists")
			default:
				return GetFileResponse{}, fmt.Errorf("s3.GetFile.service.GetFile(): %w", err)
			}
		}

		return GetFileResponse{}, fmt.Errorf("s3.GetFile.service.GetFile(): %w", err)
	}

	defer result.Body.Close()

	m := GetFileResponse{}
	fileBytes, err := io.ReadAll(result.Body)
	if err != nil {
		return GetFileResponse{}, fmt.Errorf("s3.GetFile.io.ReadAll(): %w", err)
	}
	m.FileBytes = fileBytes

	if result.ContentType != nil {
		m.ContentType = *result.ContentType
	}

	return m, nil
}

// Presign signs a key object with an expiry time
func (s Service) Presign(key string) (string, error) {
	service := s3.New(s.Session)

	// Here we just get a request with the headers needed to sign the url
	req, _ := service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})

	signedURL, err := req.Presign(time.Minute)
	if err != nil {
		return "", fmt.Errorf("s3.SignKey(): %v", err)
	}

	return signedURL, nil
}
