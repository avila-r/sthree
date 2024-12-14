package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pkg/pointer"
)

type Delete struct {
	// Indicates whether S3 Object Lock should bypass
	// Governance-mode restrictions to process this operation.
	//
	// Required permission: s3:BypassGovernanceRetention.
	BypassGovernanceRetention bool

	// The account ID of the expected bucket owner.
	ExpectedBucketOwner string

	// Key name of the object to delete.
	//
	// Required field.
	Key string

	// The concatenation of the authentication device's serial number, a space,
	// and the value that is displayed on your authentication device.
	MFA string

	// Confirms that the requester knows that they will be charged for the request.
	RequestPayer string

	// Version ID used to reference a specific version of the object.
	Version string
}

func (m *Module) Delete(key string, params ...Delete) (*s3.DeleteObjectOutput, error) {
	input := &s3.DeleteObjectInput{
		Bucket: &m.Bucket,
		Key:    &key,
	}
	if len(params) > 0 {
		input = fromDeleteParams(m.Bucket, key, params[0])
	}

	return m.Sdk.DeleteObject(input)
}

func fromDeleteParams(bucket string, key string, params Delete) *s3.DeleteObjectInput {
	return &s3.DeleteObjectInput{
		Bucket:                    pointer.NotBlank(bucket),
		BypassGovernanceRetention: &params.BypassGovernanceRetention,
		ExpectedBucketOwner:       pointer.NotBlank(params.ExpectedBucketOwner),
		Key:                       pointer.NotBlank(key),
		MFA:                       pointer.NotBlank(params.MFA),
		RequestPayer:              pointer.NotBlank(params.RequestPayer),
		VersionId:                 pointer.NotBlank(params.Version),
	}
}
