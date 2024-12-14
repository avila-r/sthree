package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type Put struct {
	Config ObjectDetails
}

func (m *Module) Put(body interface{}, params ...Put) (*s3.PutObjectOutput, error) {
	input := PutInput(m.Bucket, body, params...)

	return m.Sdk.PutObject(input)
}
