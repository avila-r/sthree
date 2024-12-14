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

func (m *Module) Get(key string, params ...Get) (*s3.GetObjectOutput, error) {
	input := GetInput(m.Bucket, key, params...)

	return m.Sdk.GetObject(input)
}

func (m *Module) Delete(key string, params ...Delete) (*s3.DeleteObjectOutput, error) {
	input := DeleteInput(m.Bucket, key, params...)

	return m.Sdk.DeleteObject(input)
}

func (m *Module) ListObjects(params ...List) (*s3.ListObjectsV2Output, error) {
	input := ListInput(m.Bucket, params...)

	return m.Sdk.ListObjectsV2(input)
}

// 'ListObjects' alias
func (m *Module) All(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}

// 'ListObjects' alias
func (m *Module) List(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}

func (m *Module) Upload(params ...Upload) (*s3manager.UploadOutput, error) {
	input := UploadInput(m.Bucket, params...)

	return m.Uploader.Upload(input)
}

func (m *Module) Put(body interface{}, params ...Put) (*s3.PutObjectOutput, error) {
	input := PutInput(m.Bucket, body, params...)

	return m.Sdk.PutObject(input)
}
