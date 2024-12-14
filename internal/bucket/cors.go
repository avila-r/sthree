package bucket

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pkg/pointer"
)

// Cors represents the configuration for Cross-Origin Resource Sharing (CORS) on an S3 bucket.
//
// It includes rules for allowed origins, headers, methods, and other options
// that control how the bucket handles cross-origin requests.
type Cors struct {
	// Bucket is the name of the bucket to apply the CORS configuration to.
	Bucket string

	// Rules is a list of CORS rules that define the allowed behaviors for cross-origin requests.
	Rules []CorsRule

	// ChecksumAlgorithm indicates the algorithm used to create the checksum for the object when using the SDK.
	ChecksumAlgorithm string

	// ExpectedBucketOwner is the account ID of the expected bucket owner.
	// This ensures that the operation only applies to a bucket owned by the specified account.
	ExpectedBucketOwner string
}

// CorsRule defines a single rule for Cross-Origin Resource Sharing (CORS).
type CorsRule struct {
	// ID is a unique identifier for the rule.
	// The value cannot exceed 255 characters.
	ID string

	// AllowedHeaders is a list of HTTP headers that are specified in the Access-Control-Request-Headers header.
	AllowedHeaders []string

	// AllowedOrigins specifies one or more origins that are permitted to access the bucket.
	// This is a required field.
	AllowedOrigins []string

	// AllowedMethods defines the HTTP methods allowed for the specified origins.
	// Valid values are GET, PUT, HEAD, POST, and DELETE.
	// This is a required field.
	AllowedMethods []string

	// ExposeHeaders is a list of response headers that are accessible to the client application.
	ExposeHeaders []string

	// MaxAgeSeconds specifies the time in seconds for which the browser should cache the preflight response.
	MaxAgeSeconds int64
}

// SetCors applies the specified CORS configuration to the bucket.
//
// This method uses the AWS S3 SDK to set the bucket's CORS configuration,
// which allows the bucket to handle cross-origin requests.
//
// @param c The Cors configuration object containing the bucket name and rules.
// @return An error if the CORS configuration cannot be applied, or nil if successful.
func (m *Module) SetCors(c Cors) error {
	_, err := m.Sdk.PutBucketCors(CorsInput(c))
	return err
}

// CorsInput converts a `Cors` struct into an AWS SDK-compatible `*s3.PutBucketCorsInput` object.
//
// This function prepares the input required to apply CORS configuration to an S3 bucket.
// It maps the Cors object to AWS's CORS rule format, ensuring compatibility with the S3 SDK.
//
// @param c The Cors struct containing configuration details.
// @return A pointer to `s3.PutBucketCorsInput`, which can be used with the AWS S3 SDK.
func CorsInput(c Cors) *s3.PutBucketCorsInput {
	rules := []*s3.CORSRule{}
	for _, rule := range c.Rules {
		rules = append(rules, &s3.CORSRule{
			ID:             pointer.NotBlank(rule.ID),
			AllowedHeaders: pointer.InSlice(rule.AllowedHeaders...),
			AllowedMethods: pointer.InSlice(rule.AllowedMethods...),
			AllowedOrigins: pointer.InSlice(rule.AllowedOrigins...),
			ExposeHeaders:  pointer.InSlice(rule.ExposeHeaders...),
			MaxAgeSeconds:  pointer.NotZero(rule.MaxAgeSeconds),
		})
	}

	return &s3.PutBucketCorsInput{
		Bucket: pointer.NotBlank(c.Bucket),
		CORSConfiguration: &s3.CORSConfiguration{
			CORSRules: rules,
		},
		ChecksumAlgorithm:   pointer.NotBlank(c.ChecksumAlgorithm),
		ExpectedBucketOwner: pointer.NotBlank(c.ExpectedBucketOwner),
	}
}
