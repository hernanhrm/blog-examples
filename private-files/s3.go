package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Permissions
const (
	// PublicRead is used to set the permission of a file to public read
	PublicRead = "public-read"
	// Private is used to set the permission of a file to private
	Private = "private"
)

// Service is the use case of the storage s3 service
type Service struct {
	Bucket  string
	Session *session.Session
	S3      *s3.S3
	Signer  *v4.Signer
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
		S3:      s3.New(awsSession),
		Signer:  v4.NewSigner(awsSession.Config.Credentials),
	}, nil
}

// Upload takes a file and uploads it to the s3 bucket
func (s Service) Upload(fileBytes []byte, contentType, path string, isPublic bool) error {
	permission := PublicRead
	if !isPublic {
		permission = Private
	}

	_, err := s.S3.PutObject(&s3.PutObjectInput{
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

// GetFile returns a file from the s3 bucket
func (s Service) GetFile(filename string) (GetFileResponse, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(filename),
	}

	result, err := s.S3.GetObject(input)
	if err != nil {
		if aswErr, ok := err.(awserr.Error); ok {
			switch aswErr.Code() {
			case s3.ErrCodeNoSuchKey:
				return GetFileResponse{}, fmt.Errorf("s3.GetFile: file does not exists")
			default:
				return GetFileResponse{}, fmt.Errorf("s3.GetFile: %w", err)
			}
		}

		return GetFileResponse{}, fmt.Errorf("s3.GetFile: %w", err)
	}
	defer result.Body.Close()

	// we need the file content type and bytes to return the file with echo (web framework)
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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.Bucket, key), nil)
	if err != nil {
		return "", fmt.Errorf("s3.Presign.http.NewRequest(): %w", err)
	}

	// we sign the request with the signTime and a time expiration of 1 day
	_, err = s.Signer.Presign(req, nil, "s3", *s.Session.Config.Region, time.Hour*24, getSignTime())
	if err != nil {
		return "", fmt.Errorf("s3.Presign.v4.NewSigner(): %w", err)
	}

	return req.URL.String(), nil
}

func (s Service) PresignV2(key string) (string, error) {
	req, _ := s.S3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	})

	signedURL, err := req.Presign(time.Minute)
	if err != nil {
		return "", fmt.Errorf("s3.SignKey(): %v", err)
	}

	return signedURL, nil
}

// getSignTime returns a truncated time to day precision
// with this precision, will get the same signature during the day
func getSignTime() time.Time {
	return time.Now().Truncate(time.Hour * 24).UTC()
}
