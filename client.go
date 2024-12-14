package sthree

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/internal/bucket"
	"github.com/avila-r/sthree/internal/buckets"
	"github.com/avila-r/sthree/internal/objects"
	"github.com/avila-r/sthree/internal/requests"
)

// Sthree represents the core struct for managing interactions with AWS S3.
// It encapsulates the AWS SDK, and modules for performing operations on buckets and requests.
type Sthree struct {
	// @param Provider: The AWS configuration provider (e.g., session or custom provider).
	Provider client.ConfigProvider

	// @param Sdk: The AWS S3 SDK instance used to interact with the S3 service.
	Sdk *s3.S3

	// @param Buckets: Module for performing bucket-level operations on S3.
	Buckets *buckets.Module

	// @param Requests: Module for handling various S3 requests.
	Requests *requests.Module
}

// Connect establishes a connection to AWS S3 using the provided ConfigProvider and optional configurations.
// It creates a new `Sthree` instance with necessary dependencies for interacting with AWS S3.
//
// @param The AWS configuration provider used to create the S3 client.
// @param Optional configuration settings to customize the AWS connection.
// @return An instance of `Sthree` initialized with the AWS SDK and required modules.
func Connect(provider client.ConfigProvider, cfgs ...Config) *Sthree {
	s3 := s3.New(provider, unwrapConfig(cfgs...)...)

	return &Sthree{
		Provider: provider,
		Sdk:      s3,
		Buckets: &buckets.Module{
			Sdk: s3,
		},
		Requests: &requests.Module{
			Sdk: s3,
		},
	}
}

// Session initializes a new session for AWS S3 and returns an instance of `Sthree`.
// It automatically creates a session with the provided configuration and AWS provider.
//
// @param Optional configuration settings to customize the AWS session.
// @return A pointer to an initialized `Sthree` instance or an error if the session creation fails.
func Session(cfgs ...Config) (*Sthree, error) {
	sess, err := session.NewSession(unwrapConfig(cfgs...)...)

	if err != nil {
		return nil, err
	}

	return Connect(sess), nil
}

// New is an alias for `Connect` and creates a new connection to AWS S3 with the provided ConfigProvider
// and optional configuration settings.
//
// @param The AWS configuration provider used to create the S3 client.
// @param Optional configuration settings to customize the AWS connection.
// @return An instance of `Sthree` initialized with the AWS SDK and required modules.
func New(provider client.ConfigProvider, cfgs ...Config) *Sthree {
	return Connect(provider, cfgs...)
}

// Client is an alias for `Connect` that creates a new connection to AWS S3 with the provided ConfigProvider
// and optional configuration settings.
//
// @param The AWS configuration provider used to create the S3 client.
// @param Optional configuration settings to customize the AWS connection.
// @return An instance of `Sthree` initialized with the AWS SDK and required modules.
func Client(provider client.ConfigProvider, cfgs ...Config) *Sthree {
	return Connect(provider, cfgs...)
}

// NewClient is an alias for `Connect` that creates a new connection to AWS S3 with the provided ConfigProvider
// and optional configuration settings.
//
// @param The AWS configuration provider used to create the S3 client.
// @param Optional configuration settings to customize the AWS connection.
// @return An instance of `Sthree` initialized with the AWS SDK and required modules.
func NewClient(provider client.ConfigProvider, cfgs ...Config) *Sthree {
	return Connect(provider, cfgs...)
}

// Bucket returns an instance of the `objects.Module` for the specified bucket.
//
// This method is used for operations involving objects within a bucket,
// such as uploading, downloading, listing, and deleting objects.
//
// @param name The name of the bucket to associate with the returned module.
// @return An instance of `objects.Module` configured with the bucket name and the associated SDK.
func (m *Sthree) Bucket(bucket string) *objects.Module {
	return &objects.Module{
		Bucket: bucket,
		Sdk:    m.Sdk,
	}
}

// From returns an instance of the `objects.Module` for the specified bucket.
//
// This method is used for operations involving objects within a bucket,
// such as uploading, downloading, listing, and deleting objects.
//
// @param name The name of the bucket to associate with the returned module.
// @return An instance of `objects.Module` configured with the bucket name and the associated SDK.
func (m *Sthree) From(bucket string) *objects.Module {
	return m.Bucket(bucket)
}

// In returns an instance of the `objects.Module` for the specified bucket.
//
// This method is used for operations involving objects within a bucket,
// such as uploading, downloading, listing, and deleting objects.
//
// @param name The name of the bucket to associate with the returned module.
// @return An instance of `objects.Module` configured with the bucket name and the associated SDK.
func (m *Sthree) In(bucket string) *objects.Module {
	return m.Bucket(bucket)
}

// For returns an instance of the `bucket.Module` for the specified bucket.
//
// This method is used to interact with bucket-specific operations,
// providing an interface to access configurations related to the bucket.
//
// @param name The name of the bucket to associate with the returned module.
// @return An instance of `bucket.Module` configured with the bucket name and the associated SDK.
func (m *Sthree) For(name string) *bucket.Module {
	return &bucket.Module{
		Bucket: name,
		Sdk:    m.Sdk,
	}
}

// unwrapConfig extracts the AWS configuration settings from the provided Config objects and
// returns them as a slice of `*aws.Config`. If no configuration is provided, it returns an empty slice.
//
// @param Optional configuration objects to customize the AWS connection.
// @return A slice of `*aws.Config` that can be used to create the AWS S3 client.
func unwrapConfig(c ...Config) []*aws.Config {
	if len(c) > 0 {
		return []*aws.Config{
			c[0].ToAWSConfig(),
		}
	}

	// Empty
	return []*aws.Config{}
}
