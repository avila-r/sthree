package buckets

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pkg/pointer"
)

// Bucket represents the configuration options for an S3 bucket.
//
// This struct encapsulates various attributes and permissions
// that can be applied when creating an S3 bucket.
type Bucket struct {
	// The type of bucket.
	// Optional field.
	Type string

	// The canned ACL to apply to the bucket.
	// Not supported for directory buckets.
	// Optional field.
	ACL string

	// Grants the specified grantee full control over the bucket (read, write, read ACP, write ACP permissions).
	// This functionality is not supported for directory buckets.
	GrantFullControl string

	// Grants the specified grantee permission to list the objects in the bucket.
	// This functionality is not supported for directory buckets.
	GrantRead string

	// Grants the specified grantee permission to read the bucket ACL.
	// This functionality is not supported for directory buckets.
	GrantReadACP string

	// Grants the specified grantee permission to create new objects in the bucket.
	// Also allows deletions and overwrites for existing objects if the grantee is the bucket or object owner.
	// This functionality is not supported for directory buckets.
	GrantWrite string

	// Grants the specified grantee permission to write the ACL for the bucket.
	// This functionality is not supported for directory buckets.
	GrantWriteACP string

	// Specifies whether to enable S3 Object Lock for the bucket.
	// This functionality is not supported for directory buckets.
	ObjectLockEnabledForBucket bool

	// Specifies the ownership control for objects in the bucket.
	ObjectOwnership string

	// Specifies the data redundancy settings for the bucket.
	// This is only supported by directory buckets.
	DataRedundancy string

	// The name of the location where the bucket will be created.
	// For directory buckets, this specifies the AZ ID of the availability zone (e.g., "usw2-az1").
	LocationName string

	// Specifies the type of location where the bucket will be created.
	LocationType string

	// Specifies the AWS region where the bucket will be created.
	LocationConstraint string
}

// Create creates a new S3 bucket with the specified name and optional parameters.
//
// If no parameters are provided, the bucket will be created with default settings.
//
// @param name The name of the bucket to create.
// @param params Optional configuration parameters for the bucket.
// @return A pointer to an s3.CreateBucketOutput containing details of the created bucket.
// @return An error if the bucket creation fails.
func (m *Module) Create(name string, params ...Bucket) (*s3.CreateBucketOutput, error) {
	input := &s3.CreateBucketInput{
		Bucket: &name,
	}
	if len(params) > 0 {
		input = fromNewParams(name, params[0])
	}

	return m.Sdk.CreateBucket(input)
}

// New is an alias for Create, providing an alternative method to create a new bucket.
//
// @param name The name of the bucket to create.
// @param params Optional configuration parameters for the bucket.
// @return A pointer to an s3.CreateBucketOutput containing details of the created bucket.
// @return An error if the bucket creation fails.
func (m *Module) New(name string, params ...Bucket) (*s3.CreateBucketOutput, error) {
	return m.Create(name, params...)
}

// fromNewParams converts the provided Bucket configuration into an s3.CreateBucketInput.
//
// This function maps the custom Bucket fields to the appropriate AWS SDK structures,
// ensuring that all fields are properly formatted.
//
// @param name The name of the bucket to create.
// @param params The custom Bucket configuration.
// @return A pointer to an s3.CreateBucketInput containing the transformed configuration.
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
