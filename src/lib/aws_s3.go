package lib

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	Env *Env
	Log *log.Logger
}

func NewS3Service(
	env *Env,
	log *log.Logger,
) *S3Service {
	return &S3Service{
		env,
		log,
	}
}

func (s3Service S3Service) UploadPhoto() {
	session, err := session.NewSession(&aws.Config{})
	if err != nil {
		s3Service.Log.Fatal(err)
	}

	// Upload Files
	err = s3Service.uploadFile(session, "test.png")
	if err != nil {
		s3Service.Log.Fatal(err)
	}
}

func (s3Service S3Service) uploadFile(session *session.Session, uploadFileDir string) error {

	upFile, err := os.Open(uploadFileDir)
	if err != nil {
		return err
	}
	defer upFile.Close()

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s3Service.Env.AWSS3Bucket),
		Key:                  aws.String(uploadFileDir),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}
