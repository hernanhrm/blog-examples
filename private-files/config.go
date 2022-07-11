package main

// Config contains the configuration read from the .env file
type Config struct {
	S3AccessKey    string `json:"s3_access_key"`
	S3SecretKey    string `json:"s3_secret_key"`
	S3BucketName   string `json:"s3_bucket_name"`
	S3BucketRegion string `json:"s3_bucket_region"`
}
