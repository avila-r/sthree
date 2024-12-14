package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type Delete struct {
	// The name of the bucket containing the object.
	Bucket string

	// Key name of the object to delete.
	Key string

	// Indicates whether S3 Object Lock should bypass
	// Governance-mode restrictions to process this operation.
	//
	// Required permission: s3:BypassGovernanceRetention.
	BypassGovernanceRetention bool

	// The account ID of the expected bucket owner.
	ExpectedBucketOwner string

	// The concatenation of the authentication device's serial number, a space,
	// and the value that is displayed on your authentication device.
	MFA string

	// Confirms that the requester knows that they will be charged for the request.
	RequestPayer string

	// Version ID used to reference a specific version of the object.
	Version string
}

func (m *Module) Delete(key string, params ...Delete) (*s3.DeleteObjectOutput, error) {
	input := DeleteInput(m.Bucket, key, params...)

	return m.Sdk.DeleteObject(input)
}
