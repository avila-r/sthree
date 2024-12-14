package bucket

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/internal/objects"
)

// bucket.Module represents the S3 bucket configuration and provides methods for interacting with it.
//
// This struct acts as a wrapper for the AWS S3 SDK and is used to manage bucket-related operations.
type Module struct {
	// Bucket is the name of the S3 bucket associated with this module.
	Bucket string

	// Sdk is the AWS S3 SDK client used to interact with the S3 bucket.
	Sdk *s3.S3
}

// Objects creates and returns a new `objects.Module` instance.
//
// The returned module provides functionality for managing objects within the S3 bucket.
// This method ensures that the `objects.Module` is initialized with the same bucket name
// and SDK client as the current `bucket.Module`.
//
// @return A pointer to an `objects.Module` instance configured for the current S3 bucket.
func (m *Module) Objects() *objects.Module {
	return &objects.Module{
		Bucket: m.Bucket,
		Sdk:    m.Sdk,
	}
}
