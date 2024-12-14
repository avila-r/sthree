package buckets

import "github.com/aws/aws-sdk-go/service/s3"

// List retrieves a list of all buckets owned by the sender.
//
// Required permissions:
// - s3:ListAllMyBuckets
//
// @return A pointer to an s3.ListBucketsOutput containing details of all the buckets.
// @return An error if the operation fails.
func (m *Module) List() (*s3.ListBucketsOutput, error) {
	return m.Sdk.ListBuckets(nil)
}

// All is an alias for the List method, providing an alternative way to
// retrieve the list of all buckets owned by the sender.
//
// Required permissions:
// - s3:ListAllMyBuckets
//
// @return A pointer to an s3.ListBucketsOutput containing details of all the buckets.
// @return An error if the operation fails.
func (m *Module) All() (*s3.ListBucketsOutput, error) {
	return m.List()
}
