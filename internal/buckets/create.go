package buckets

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pkg/pointer"
)

type Bucket struct {
	// The type of bucket.
	//
	// Optional field.
	Type string

	// The canned ACL to apply to the bucket.
	//
	// Not supported for directory buckets.
	//
	// Optional field.
	ACL string

	// Allows grantee the read, write, read ACP,
	// and write ACP permissions on the bucket.
	//
	// This functionality is not supported
	// for directory buckets.
	GrantFullControl string

	// Allows grantee to list the objects in the bucket.
	//
	// This functionality is not supported for directory buckets.
	GrantRead string

	// Allows grantee to read the bucket ACL.
	//
	// This functionality is not supported for directory buckets.
	GrantReadACP string

	// Allows grantee to create new objects in the bucket.
	//
	// For the bucket and object owners of existing objects, also allows deletions
	// and overwrites of those objects.
	//
	// This functionality is not supported for directory buckets.
	GrantWrite string

	// Allows grantee to write the ACL for the applicable bucket.
	//
	// This functionality is not supported for directory buckets.
	GrantWriteACP string

	// Specifies whether you want S3 Object Lock
	// to be enabled for the new bucket.
	//
	// This functionality is not supported
	// for directory buckets.
	ObjectLockEnabledForBucket bool

	// The container element for object ownership
	// for a bucket's ownership controls.
	ObjectOwnership string

	// Specifies the information about
	// the bucket that will be created.
	//
	// This functionality is only
	// supported by directory buckets.
	DataRedundancy string

	// The name of the location where
	// the bucket will be created.
	//
	// For directory buckets, the name of the location
	// is the AZ ID of the Availability
	//
	// Zone where the bucket will be created.
	// An example AZ ID value is usw2-az1.
	LocationName string

	// The type of location where the bucket will be created.
	LocationType string

	// Specifies the Region where the bucket will be created.
	LocationConstraint string
}

func (m *Module) Create(name string, params ...Bucket) (*s3.CreateBucketOutput, error) {
	input := &s3.CreateBucketInput{
		Bucket: &name,
	}
	if len(params) > 0 {
		input = fromNewParams(name, params[0])
	}

	return m.Sdk.CreateBucket(input)
}

func (m *Module) New(name string, params ...Bucket) (*s3.CreateBucketOutput, error) {
	return m.Create(name, params...)
}

func fromNewParams(name string, params Bucket) *s3.CreateBucketInput {
	return &s3.CreateBucketInput{
		Bucket: &name,
		ACL:    pointer.NotBlank(params.ACL),

		GrantFullControl: pointer.NotBlank(params.GrantFullControl),
		GrantRead:        pointer.NotBlank(params.GrantRead),
		GrantReadACP:     pointer.NotBlank(params.GrantReadACP),
		GrantWrite:       pointer.NotBlank(params.GrantWrite),
		GrantWriteACP:    pointer.NotBlank(params.GrantWriteACP),

		ObjectOwnership:            pointer.NotBlank(params.ObjectOwnership),
		ObjectLockEnabledForBucket: &params.ObjectLockEnabledForBucket,

		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: pointer.NotBlank(params.LocationConstraint),

			Bucket: &s3.BucketInfo{
				Type:           pointer.NotBlank(params.Type),
				DataRedundancy: pointer.NotBlank(params.DataRedundancy),
			},

			Location: &s3.LocationInfo{
				Name: pointer.NotBlank(params.LocationName),
				Type: pointer.NotBlank(params.LocationType),
			},
		},
	}
}
