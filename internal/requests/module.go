package requests

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

// Module is a wrapper around the AWS S3 SDK client, allowing operations
// to be performed using the AWS SDK methods for S3 services.
type Module struct {
	Sdk *s3.S3 // The S3 client to interact with S3 services.
}
