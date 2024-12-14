package objects

// Put represents the parameters required to upload an object to an S3 bucket.
// This struct is used in the PutObject API call to specify the object to be uploaded
// and additional configuration settings.
//
// Fields:
//   - Bucket: The name of the S3 bucket where the object will be uploaded.
//   - Body: The body (content) of the object to upload. This can be of any type
//     that can be serialized to be uploaded to S3 (e.g., a file, a string, etc.).
//   - Config: Additional details or configuration specific to the object upload,
//     such as metadata, ACL (Access Control List), content type, etc.
type Put struct {
	// Bucket specifies the name of the S3 bucket where the object will be stored.
	Bucket string

	// Body contains the content of the object to be uploaded to S3.
	// This can be any type that is serializable into an S3-compatible format.
	Body interface{}

	// Config contains additional details for uploading the object, such as metadata
	// or access control settings. This field is of type ObjectDetails, which contains
	// settings related to the object configuration.
	Config ObjectDetails
}
