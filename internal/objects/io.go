package objects

import (
	"bytes"
	"encoding/json"
	"io"
)

// Upload represents an object for uploading data to S3 with the provided readable body and object details.
type Upload struct {
	// Body holds the readable body payload to send to S3.
	Body io.Reader
	// ObjectDetails holds additional details of the object being uploaded.
	ObjectDetails
}

// ToReadSeeker converts a given object into an io.ReadSeeker, suitable for uploading to S3.
// It marshals the object into JSON and returns a ReadSeeker for the resulting byte array.
// If the object cannot be marshaled, it returns an empty byte reader.
//
// @param t The object to be converted into a ReadSeeker.
// @return An io.ReadSeeker containing the marshaled JSON of the object, or an empty ReadSeeker in case of error.
func ToReadSeeker(t any) io.ReadSeeker {
	// Marshal the object into JSON
	json, err := json.Marshal(t)
	if err != nil {
		// If marshalling fails, return an empty ReadSeeker
		return bytes.NewReader([]byte{})
	}

	// Return the ReadSeeker for the JSON byte array
	return bytes.NewReader(json)
}
