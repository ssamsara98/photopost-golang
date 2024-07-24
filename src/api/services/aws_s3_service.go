package services

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/ssamsara98/photopost-golang/src/lib"
)

type S3Service struct {
	Env    *lib.Env
	logger *lib.Logger
}

func NewS3Service(
	env *lib.Env,
	logger *lib.Logger,
) *S3Service {
	return &S3Service{
		env,
		logger,
	}
}

func (s3Service S3Service) UploadPhoto(image *multipart.FileHeader) (*manager.UploadOutput, error) {
	filepath, err := uuid.NewRandom()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	id, err := gonanoid.New(32)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	file, err := image.Open()
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	defer file.Close()

	creds := credentials.NewStaticCredentialsProvider(s3Service.Env.AWSAccessKeyID, s3Service.Env.AWSSecretAccessKey, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(s3Service.Env.AWSRegion))
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	fileExt := path.Ext(image.Filename)
	keypath := fmt.Sprintf("%s/%s%s", filepath.String(), id, fileExt)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(s3Service.Env.AWSS3Bucket),
		Key:           aws.String(fmt.Sprintf("img/photopost/%s", keypath)),
		Body:          file,
		ACL:           types.ObjectCannedACLPublicRead,
		ContentType:   aws.String(image.Header.Get("Content-Type")),
		ContentLength: &image.Size,
		// ContentDisposition:   aws.String("attachment"),
		// ServerSideEncryption: types.ServerSideEncryptionAes256,
	})
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	return result, nil
}
