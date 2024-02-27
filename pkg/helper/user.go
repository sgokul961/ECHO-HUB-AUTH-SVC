package helper

import (
	"context"
	"mime/multipart"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func AddImagesToS3Bucket(image *multipart.FileHeader) (string, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-south-1"))
	if err != nil {
		return "", err
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	uploader := manager.NewUploader(s3Client)

	file, err := image.Open()

	if err != nil {
		return "", err
	}
	defer file.Close()

	upload, err1 := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("time-peace"),
		Key:    aws.String(image.Filename),
		Body:   file,
		ACL:    "public-read",
	})
	if err1 != nil {
		
		return "", err1
	}
	return upload.Location, nil

}

func IsValidEmail(email string) bool {

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}
func IsValidPhoneNumber(phoneNumber string) bool {
	regex := `^[0-9]{10}$`
	match, _ := regexp.MatchString(regex, phoneNumber)
	return match
}
