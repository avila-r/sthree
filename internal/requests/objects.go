package requests

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/internal/objects"
)

// GetObject initiates an S3 GetObject request using the provided 'Get' object from the 'objects' package.
// This method returns the AWS SDK's request object and the resulting 'GetObject' output.
//
// Parameters:
//   - from: An 'objects.Get' struct containing the S3 bucket name, key (object identifier),
//     and any additional parameters necessary for the 'GetObject' operation.
//
// Returns:
// - A request object for the S3 GetObject operation.
// - A 'GetObjectOutput' containing the details of the retrieved object.
func (m *Module) GetObject(from objects.Get) (*request.Request, *s3.GetObjectOutput) {
	input := objects.GetInput(from.Bucket, from.Key, from)

	return m.Sdk.GetObjectRequest(input)
}

// DeleteObject initiates an S3 DeleteObject request using the provided 'Delete' object from the 'objects' package.
// This method returns the AWS SDK's request object and the resulting 'DeleteObject' output.
//
// Parameters:
//   - from: An 'objects.Delete' struct containing the S3 bucket name, key (object identifier),
//     and any additional parameters necessary for the 'DeleteObject' operation.
//
// Returns:
// - A request object for the S3 DeleteObject operation.
// - A 'DeleteObjectOutput' containing the details of the deleted object.
func (m *Module) DeleteObject(from objects.Delete) (*request.Request, *s3.DeleteObjectOutput) {
	input := objects.DeleteInput(from.Bucket, from.Key, from)

	return m.Sdk.DeleteObjectRequest(input)
}

// ListObjects initiates an S3 ListObjectsV2 request using the provided 'List' object from the 'objects' package.
// This method returns the AWS SDK's request object and the resulting 'ListObjectsV2' output.
//
// Parameters:
//   - from: An 'objects.List' struct containing the S3 bucket name and any additional parameters
//     necessary for the 'ListObjectsV2' operation.
//
// Returns:
// - A request object for the S3 ListObjectsV2 operation.
// - A 'ListObjectsV2Output' containing a list of objects from the S3 bucket.
func (m *Module) ListObjects(from objects.List) (*request.Request, *s3.ListObjectsV2Output) {
	input := objects.ListInput(from.Bucket, from)

	return m.Sdk.ListObjectsV2Request(input)
}

// PutObject initiates an S3 PutObject request using the provided 'Put' object from the 'objects' package.
// This method returns the AWS SDK's request object and the resulting 'PutObject' output.
//
// Parameters:
//   - from: An 'objects.Put' struct containing the S3 bucket name, the object body, and any additional
//     configuration details necessary for the 'PutObject' operation.
//
// Returns:
// - A request object for the S3 PutObject operation.
// - A 'PutObjectOutput' containing the details of the uploaded object.
func (m *Module) PutObject(from objects.Put) (*request.Request, *s3.PutObjectOutput) {
	input := objects.PutInput(from.Bucket, from.Body, from)

	return m.Sdk.PutObjectRequest(input)
}
