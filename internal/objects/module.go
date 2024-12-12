package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Module struct {
	Bucket   string
	Sdk      *s3.S3
	Uploader s3manager.Uploader
}
