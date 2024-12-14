package objects

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Upload struct {
	// The readable body payload to send to S3.
	Body io.Reader
	ObjectDetails
}

func ToReadSeeker(t any) io.ReadSeeker {
	json, err := json.Marshal(t)
	if err != nil {
		return bytes.NewReader([]byte{})
	}

	return bytes.NewReader(json)
}

func (m *Module) Upload(params ...Upload) (*s3manager.UploadOutput, error) {
	input := UploadInput(m.Bucket, params...)

	return m.Uploader.Upload(input)
}
