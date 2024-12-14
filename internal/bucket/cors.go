package bucket

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pkg/pointer"
)

type Cors struct {
	Bucket string

	Rules []CorsRule

	// Indicates the algorithm used to create the checksum for the object when you
	// use the SDK.
	ChecksumAlgorithm string

	// The account ID of the expected bucket owner.
	ExpectedBucketOwner string
}

type CorsRule struct {
	// Unique identifier for the rule. The value cannot be longer than 255 characters.
	ID string

	// Headers that are specified in the Access-Control-Request-Headers header.
	AllowedHeaders []string

	// One or more origins you want customers to be able to access the bucket from.
	//
	// Required field.
	AllowedOrigins []string

	// An HTTP method that you allow the origin to execute. Valid values are GET,
	// PUT, HEAD, POST, and DELETE.
	//
	// Required field.
	AllowedMethods []string

	// One or more headers in the response that you want customers to be able to
	// access from their applications.
	ExposeHeaders []string

	// The time in seconds that your browser is to cache the preflight response
	// for the specified resource.
	MaxAgeSeconds int64
}

func (m *Module) SetCors(c Cors) error {
	_, err := m.Sdk.PutBucketCors(CorsInput(c))
	return err
}

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
