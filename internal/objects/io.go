package objects

import (
	"bytes"
	"encoding/json"
	"io"
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
