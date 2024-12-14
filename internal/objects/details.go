package objects

import (
	"time"
)

type ObjectDetails struct {
	// The canned ACL to apply to the object.
	//
	// This functionality is not supported for
	// directory buckets or Amazon S3 on Outposts.
	ACL string

	// The bucket name to which the PUT action was initiated.
	//
	// Required field.
	Bucket string

	// Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption
	// with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
	//
	// This functionality is not supported for directory buckets.
	BucketKeyEnabled bool

	// Can be used to specify caching behavior along the request/reply chain.
	CacheControl string

	// Indicates the algorithm used to create the
	// checksum for the object when you use the SDK.
	ChecksumAlgorithm string

	// This header can be used as a data integrity
	// check to verify that the data received is
	// the same data that was originally sent.
	ChecksumCRC32 string

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent.
	ChecksumCRC32C string

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent.
	ChecksumSHA1 string

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent.
	ChecksumSHA256 string

	// Specifies presentational information for the object.
	ContentDisposition string

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field.
	ContentEncoding string

	// The language the content is in.
	ContentLanguage string

	// The base64-encoded 128-bit MD5 digest of the message
	// without the headers) according to RFC 1864.
	ContentMD5 string

	// A standard MIME type describing the format of the contents.
	ContentType string

	// The account ID of the expected bucket owner.
	ExpectedBucketOwner string

	// The date and time at which the object is no longer cacheable.
	Expires time.Time

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	//
	// This functionality is not supported for
	// directory buckets or Amazon S3 on Outposts.
	GrantFullControl string

	// Allows grantee to read the object data and its metadata.
	//
	// This functionality is not supported for
	// directory buckets or Amazon S3 on Outposts.
	GrantRead string

	// Allows grantee to read the object ACL.
	//
	// This functionality is not supported for
	// directory buckets or Amazon S3 on Outposts.
	GrantReadACP string

	// Allows grantee to write the ACL for the applicable object.
	//
	// This functionality is not supported for
	// directory buckets or Amazon S3 on Outposts.
	GrantWriteACP string

	// Object key for which the PUT action was initiated.
	//
	// Required field.
	Key string

	// A map of metadata to store with the object in S3.
	Metadata map[string]string

	// Specifies whether a legal hold will be applied to this object.
	//
	// This functionality is not supported for directory buckets.
	ObjectLockLegalHoldStatus string

	// The Object Lock mode that you want to apply to this object.
	//
	// This functionality is not supported for directory buckets.
	ObjectLockMode string

	// The date and time when you want this object's Object Lock to expire. Must
	// be formatted as a timestamp parameter.
	//
	// This functionality is not supported for directory buckets.
	ObjectLockRetainUntilDate time.Time

	// Confirms that the requester knows that they will be charged for the request.
	//
	// This functionality is not supported for directory buckets.
	RequestPayer string

	// Specifies the algorithm to use when encrypting the object (for example, AES256).
	//
	// This functionality is not supported for directory buckets.
	SSECustomerAlgorithm string

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data.
	//
	// This functionality is not supported for directory buckets.
	SSECustomerKey string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	//
	// This functionality is not supported for directory buckets.
	SSECustomerKeyMD5       string
	SSEKMSEncryptionContext string
	SSEKMSKeyId             string

	// The server-side encryption algorithm that was used when you store this object
	// in Amazon S3 (for example, AES256, aws:kms, aws:kms:dsse).
	ServerSideEncryption string

	StorageClass string

	// The tag-set for the object. The tag-set must be encoded as URL Query parameters.
	// (For example, "Key1=Value1")
	//
	// This functionality is not supported for directory buckets.
	Tagging string

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata. For information about object
	// metadata, see Object Key and Metadata (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html)
	// in the Amazon S3 User Guide.
	WebsiteRedirectLocation string
}
