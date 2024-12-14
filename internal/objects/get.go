package objects

import (
	"time"
)

type Get struct {
	// The name of the bucket containing the object.
	Bucket string

	// Key name of the object to delete.
	Key string

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
