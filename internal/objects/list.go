package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

type List struct {
	// The name of the bucket containing the object.
	Bucket string

	// Key name of the object to delete.
	Key string

	// ContinuationToken indicates to Amazon S3 that the list is being continued
	// on this bucket with a token. ContinuationToken is obfuscated and is not a
	// real key. You can use this ContinuationToken for pagination of the list results.
	ContinuationToken string

	// A delimiter is a character that you use to group keys.
	Delimiter string

	// Encoding type used by Amazon S3 to encode object keys in the response.
	EncodingType string

	// The account ID of the expected bucket owner.
	ExpectedBucketOwner string

	// The owner field is not present in V2 mode by default. If you want to
	// return the owner field with each key in the result, then set the FetchOwner
	// field to true.
	FetchOwner bool

	// Sets the maximum number of keys
	// returned in the response.
	MaxKeys int64

	// Specifies the optional fields that
	// you want returned in the response.
	//
	// Fields that you do not specify are not returned.
	OptionalObjectAttributes []string

	// Limits the response to keys that
	// begin with the specified prefix.
	Prefix string

	// Confirms that the requester knows that
	// she or he will be charged for the
	// list objects request in V2 style.
	RequestPayer string

	// StartAfter is where you want
	// Amazon S3 to start listing from.
	StartAfter string
}

func (m *Module) ListObjects(params ...List) (*s3.ListObjectsV2Output, error) {
	input := ListInput(m.Bucket, params...)

	return m.Sdk.ListObjectsV2(input)
}

// 'ListObjects' alias
func (m *Module) All(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}

// 'ListObjects' alias
func (m *Module) List(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}
