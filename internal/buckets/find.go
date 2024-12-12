package buckets

import "github.com/aws/aws-sdk-go/service/s3"

func (m *Module) List() (*s3.ListBucketsOutput, error) {
	return m.Sdk.ListBuckets(nil)
}

func (m *Module) All() (*s3.ListBucketsOutput, error) {
	return m.List()
}
