package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// objects.Module represents a service for interacting with an S3 bucket,
// providing methods to get, delete, list, and upload objects.
type Module struct {
	// Bucket is the name of the S3 bucket.
	Bucket string
	// Sdk is the AWS SDK client for interacting with S3.
	Sdk *s3.S3
	// Uploader is used to upload objects to S3.
	Uploader s3manager.Uploader
}

// Get retrieves an object from the S3 bucket by key.
// It takes additional parameters for customizing the request.
//
// @param key The key (filename) of the object to retrieve from S3.
// @param params Optional additional parameters for customizing the request (e.g., range, version).
// @return A pointer to the GetObjectOutput containing the retrieved object, or an error if the operation fails.
func (m *Module) Get(key string, params ...Get) (*s3.GetObjectOutput, error) {
	input := GetInput(m.Bucket, key, params...)

	return m.Sdk.GetObject(input)
}

// Delete deletes an object from the S3 bucket by key.
// It takes additional parameters for customizing the request.
//
// @param key The key (filename) of the object to delete from S3.
// @param params Optional additional parameters for customizing the request (e.g., version).
// @return A pointer to the DeleteObjectOutput indicating the result of the deletion, or an error.
func (m *Module) Delete(key string, params ...Delete) (*s3.DeleteObjectOutput, error) {
	input := DeleteInput(m.Bucket, key, params...)

	return m.Sdk.DeleteObject(input)
}

// ListObjects lists the objects in the S3 bucket.
// It takes additional parameters for customizing the listing request.
//
// @param params Optional parameters for customizing the listing (e.g., prefix, delimiter, max keys).
// @return A pointer to the ListObjectsV2Output containing the list of objects, or an error.
func (m *Module) ListObjects(params ...List) (*s3.ListObjectsV2Output, error) {
	input := ListInput(m.Bucket, params...)

	return m.Sdk.ListObjectsV2(input)
}

// All is an alias for ListObjects to retrieve all objects in the S3 bucket.
// It functions the same as ListObjects, providing an easy way to call the method.
//
// @param params Optional parameters for customizing the listing (e.g., prefix, delimiter, max keys).
// @return A pointer to the ListObjectsV2Output containing the list of objects, or an error.
func (m *Module) All(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}

// List is an alias for ListObjects to retrieve a list of objects from the S3 bucket.
// It functions the same as ListObjects, providing an alternate name for the method.
//
// @param params Optional parameters for customizing the listing (e.g., prefix, delimiter, max keys).
// @return A pointer to the ListObjectsV2Output containing the list of objects, or an error.
func (m *Module) List(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}

// Upload uploads an object to the S3 bucket.
// It takes additional parameters for customizing the upload request.
//
// @param params Optional parameters for customizing the upload (e.g., content type, ACL).
// @return A pointer to the UploadOutput indicating the result of the upload, or an error.
func (m *Module) Upload(params ...Upload) (*s3manager.UploadOutput, error) {
	input := UploadInput(m.Bucket, params...)

	return m.Uploader.Upload(input)
}

// Put uploads an object to the S3 bucket using the PutObject API.
// It takes the body of the object and additional parameters for customization.
//
// @param body The body of the object to upload.
// @param params Optional parameters for customizing the upload (e.g., content type, ACL).
// @return A pointer to the PutObjectOutput indicating the result of the upload, or an error.
func (m *Module) Put(body interface{}, params ...Put) (*s3.PutObjectOutput, error) {
	input := PutInput(m.Bucket, body, params...)

	return m.Sdk.PutObject(input)
}
