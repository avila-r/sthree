package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pkg/pointer"
)

type List struct {
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
	input := &s3.ListObjectsV2Input{
		Bucket: &m.Bucket,
	}
	if len(params) > 0 {
		input = fromListParams(m.Bucket, params[0])
	}

	return m.Sdk.ListObjectsV2(input)
}

func fromListParams(name string, params List) *s3.ListObjectsV2Input {
	input := &s3.ListObjectsV2Input{
		Bucket: &name,

		ContinuationToken:   pointer.NotBlank(params.ContinuationToken),
		Delimiter:           pointer.NotBlank(params.Delimiter),
		EncodingType:        pointer.NotBlank(params.EncodingType),
		ExpectedBucketOwner: pointer.NotBlank(params.ExpectedBucketOwner),
		FetchOwner:          &params.FetchOwner,

		Prefix:       pointer.NotBlank(params.Prefix),
		RequestPayer: pointer.NotBlank(params.RequestPayer),
		StartAfter:   pointer.NotBlank(params.StartAfter),
	}

	input.MaxKeys = pointer.Of[int64](1000)
	if params.MaxKeys > 0 {
		input.MaxKeys = &params.MaxKeys
	}

	attrs := []*string{}
	for _, attr := range params.OptionalObjectAttributes {
		attrs = append(attrs, &attr)
	}
	input.OptionalObjectAttributes = attrs

	return input
}

// 'ListObjects' alias
func (m *Module) All(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}

// 'ListObjects' alias
func (m *Module) List(params ...List) (*s3.ListObjectsV2Output, error) {
	return m.ListObjects(params...)
}
