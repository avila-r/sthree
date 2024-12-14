package objects

import (
	"time"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pkg/pointer"
)

type Get struct {
	// To retrieve the checksum, this mode must be enabled.
	ChecksumMode string

	// The account ID of the expected bucket owner.
	ExpectedBucketOwner string

	// Return the object only if its entity tag (ETag) is the same as the one specified
	// in this header; otherwise, return a 412 Precondition Failed error.
	IfMatch string

	// Return the object only if it has been modified since the specified time;
	// otherwise, return a 304 Not Modified error.
	IfModifiedSince time.Time

	// Return the object only if its entity tag (ETag) is different from the one
	// specified in this header; otherwise, return a 304 Not Modified error.
	IfNoneMatch string

	// Return the object only if it has not been modified since the specified time;
	// otherwise, return a 412 Precondition Failed error.
	IfUnmodifiedSince time.Time

	// Part number of the object being read.
	PartNumber int64

	// Downloads the specified byte range of an object.
	Range string

	// Confirms that the requester knows that they will be charged for the request.
	//
	// This functionality is not supported for directory buckets.
	RequestPayer string

	// Sets the Cache-Control header of the response.
	ResponseCacheControl string

	// Sets the Content-Disposition header of the response.
	ResponseContentDisposition string

	// Sets the Content-Encoding header of the response.
	ResponseContentEncoding string

	// Sets the Content-Language header of the response.
	ResponseContentLanguage string

	// Sets the Content-Type header of the response.
	ResponseContentType string

	// Sets the Expires header of the response.
	ResponseExpires time.Time

	// Specifies the algorithm to use when decrypting the object (for example, AES256).
	SSECustomerAlgorithm string

	// Specifies the customer-provided encryption key that you originally provided
	// for Amazon S3 to encrypt the data before storing it.
	SSECustomerKey string

	// Specifies the 128-bit MD5 digest of the customer-provided encryption key
	// according to RFC 1321. Amazon S3 uses this header for a message integrity
	// check to ensure that the encryption key was transmitted without error.
	SSECustomerKeyMD5 string

	// Version ID used to reference a specific version of the object.
	Version string
}

func (m *Module) Get(key string, params ...Get) (*s3.GetObjectOutput, error) {
	input := &s3.GetObjectInput{
		Bucket: &m.Bucket,
		Key:    &key,
	}
	if len(params) > 0 {
		input = fromGetParams(m.Bucket, key, params[0])
	}

	return m.Sdk.GetObject(input)
}

func fromGetParams(bucket string, key string, params Get) *s3.GetObjectInput {
	return &s3.GetObjectInput{
		Bucket:              pointer.NotBlank(bucket),
		ChecksumMode:        pointer.NotBlank(params.ChecksumMode),
		ExpectedBucketOwner: pointer.NotBlank(params.ExpectedBucketOwner),

		IfMatch:           pointer.NotBlank(params.IfMatch),
		IfModifiedSince:   pointer.Time(params.IfModifiedSince),
		IfNoneMatch:       pointer.NotBlank(params.IfNoneMatch),
		IfUnmodifiedSince: pointer.Time(params.IfUnmodifiedSince),

		Key:          pointer.NotBlank(key),
		PartNumber:   pointer.NotZero(params.PartNumber),
		Range:        pointer.NotBlank(params.Range),
		RequestPayer: pointer.NotBlank(params.RequestPayer),

		ResponseCacheControl:       pointer.NotBlank(params.ResponseCacheControl),
		ResponseContentDisposition: pointer.NotBlank(params.ResponseContentDisposition),
		ResponseContentEncoding:    pointer.NotBlank(params.ResponseContentEncoding),
		ResponseContentLanguage:    pointer.NotBlank(params.ResponseContentLanguage),
		ResponseContentType:        pointer.NotBlank(params.ResponseContentType),
		ResponseExpires:            pointer.Time(params.ResponseExpires),

		SSECustomerAlgorithm: pointer.NotBlank(params.SSECustomerAlgorithm),
		SSECustomerKey:       pointer.NotBlank(params.SSECustomerKey),
		SSECustomerKeyMD5:    pointer.NotBlank(params.SSECustomerKeyMD5),

		VersionId: pointer.NotBlank(params.Version),
	}
}
