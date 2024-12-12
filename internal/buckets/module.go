package buckets

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type Module struct {
	Sdk *s3.S3
}
