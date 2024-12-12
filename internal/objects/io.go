package objects

import (
	"bytes"
	"encoding/json"
	"io"
)

func ToReadSeeker(t any) io.ReadSeeker {
	json, err := json.Marshal(t)
	if err != nil {
		return bytes.NewReader([]byte{})
	}

	return bytes.NewReader(json)
}
