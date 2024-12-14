package bucket

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/internal/objects"
)

type Module struct {
	Bucket string
	Sdk    *s3.S3
}

func (m *Module) Objects() *objects.Module {
	return &objects.Module{
		Bucket: m.Bucket,
		Sdk:    m.Sdk,
	}
}
