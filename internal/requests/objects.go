package requests

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/internal/objects"
)

func (m *Module) GetObject(from objects.Get) (*request.Request, *s3.GetObjectOutput) {
	input := objects.GetInput(from.Bucket, from.Key, from)

	return m.Sdk.GetObjectRequest(input)
}

func (m *Module) DeleteObject(from objects.Delete) (*request.Request, *s3.DeleteObjectOutput) {
	input := objects.DeleteInput(from.Bucket, from.Key, from)

	return m.Sdk.DeleteObjectRequest(input)
}

func (m *Module) ListObjects(from objects.List) (*request.Request, *s3.ListObjectsV2Output) {
	input := objects.ListInput(from.Bucket, from)

	return m.Sdk.ListObjectsV2Request(input)
}

func (m *Module) PutObject(from objects.Put) (*request.Request, *s3.PutObjectOutput) {
	input := objects.PutInput(from.Bucket, from.Body, from)

	return m.Sdk.PutObjectRequest(input)
}
