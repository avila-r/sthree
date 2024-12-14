package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/avila-r/sthree/pkg/pointer"
)

// GetInput constructs an s3.GetObjectInput for retrieving an object from S3.
//
// This method accepts a bucket name and an object key, with optional parameters
// to customize the request. If no parameters are provided, default values are used.
//
// @param bucket The name of the S3 bucket containing the object.
// @param key The key of the object to retrieve.
// @param params Optional configuration parameters for the request.
// @return A pointer to an s3.GetObjectInput with the configured values.
func GetInput(bucket string, key string, params ...Get) *s3.GetObjectInput {
	cfg := Get{}
	if len(params) > 0 {
		cfg = params[0]
	}

	return &s3.GetObjectInput{
		Bucket:              pointer.NotBlank(bucket),
		ChecksumMode:        pointer.NotBlank(cfg.ChecksumMode),
		ExpectedBucketOwner: pointer.NotBlank(cfg.ExpectedBucketOwner),

		IfMatch:           pointer.NotBlank(cfg.IfMatch),
		IfModifiedSince:   pointer.Time(cfg.IfModifiedSince),
		IfNoneMatch:       pointer.NotBlank(cfg.IfNoneMatch),
		IfUnmodifiedSince: pointer.Time(cfg.IfUnmodifiedSince),

		Key:          pointer.NotBlank(key),
		PartNumber:   pointer.NotZero(cfg.PartNumber),
		Range:        pointer.NotBlank(cfg.Range),
		RequestPayer: pointer.NotBlank(cfg.RequestPayer),

		ResponseCacheControl:       pointer.NotBlank(cfg.ResponseCacheControl),
		ResponseContentDisposition: pointer.NotBlank(cfg.ResponseContentDisposition),
		ResponseContentEncoding:    pointer.NotBlank(cfg.ResponseContentEncoding),
		ResponseContentLanguage:    pointer.NotBlank(cfg.ResponseContentLanguage),
		ResponseContentType:        pointer.NotBlank(cfg.ResponseContentType),
		ResponseExpires:            pointer.Time(cfg.ResponseExpires),

		SSECustomerAlgorithm: pointer.NotBlank(cfg.SSECustomerAlgorithm),
		SSECustomerKey:       pointer.NotBlank(cfg.SSECustomerKey),
		SSECustomerKeyMD5:    pointer.NotBlank(cfg.SSECustomerKeyMD5),

		VersionId: pointer.NotBlank(cfg.Version),
	}
}

// ListInput constructs an s3.ListObjectsV2Input for listing objects in an S3 bucket.
//
// This method accepts a bucket name and optional parameters to filter the listing.
// If no parameters are provided, default values are used.
//
// @param bucket The name of the S3 bucket to list objects from.
// @param params Optional configuration parameters for the listing.
// @return A pointer to an s3.ListObjectsV2Input with the configured values.
func ListInput(bucket string, params ...List) *s3.ListObjectsV2Input {
	cfg := List{}
	if len(params) > 0 {
		cfg = params[0]
	}

	input := &s3.ListObjectsV2Input{
		Bucket: pointer.NotBlank(bucket),

		ContinuationToken:   pointer.NotBlank(cfg.ContinuationToken),
		Delimiter:           pointer.NotBlank(cfg.Delimiter),
		EncodingType:        pointer.NotBlank(cfg.EncodingType),
		ExpectedBucketOwner: pointer.NotBlank(cfg.ExpectedBucketOwner),
		FetchOwner:          pointer.NotFalse(cfg.FetchOwner),

		Prefix:       pointer.NotBlank(cfg.Prefix),
		RequestPayer: pointer.NotBlank(cfg.RequestPayer),
		StartAfter:   pointer.NotBlank(cfg.StartAfter),
	}

	input.MaxKeys = pointer.Of[int64](1000)
	if cfg.MaxKeys > 0 {
		input.MaxKeys = &cfg.MaxKeys
	}

	attrs := []*string{}
	if attributes := cfg.OptionalObjectAttributes; attributes != nil {
		for _, attr := range attributes {
			attrs = append(attrs, &attr)
		}
	}
	input.OptionalObjectAttributes = attrs

	return input
}

// PutInput constructs an s3.PutObjectInput for uploading an object to S3.
//
// This method accepts a bucket name, object body, and optional configuration parameters.
// If no parameters are provided, default values are used.
//
// @param bucket The name of the S3 bucket to upload the object to.
// @param body The body of the object to upload.
// @param params Optional configuration parameters for the upload.
// @return A pointer to an s3.PutObjectInput with the configured values.
func PutInput(bucket string, body interface{}, params ...Put) *s3.PutObjectInput {
	cfg := Put{}
	if len(params) > 0 {
		cfg = params[0]
	}

	input := &s3.PutObjectInput{
		Bucket:           pointer.NotBlank(bucket),
		ACL:              pointer.NotBlank(cfg.Config.ACL),
		Body:             ToReadSeeker(body),
		BucketKeyEnabled: pointer.NotFalse(cfg.Config.BucketKeyEnabled),
		CacheControl:     pointer.NotBlank(cfg.Config.CacheControl),

		ChecksumAlgorithm: pointer.NotBlank(cfg.Config.ChecksumAlgorithm),
		ChecksumCRC32:     pointer.NotBlank(cfg.Config.ChecksumCRC32),
		ChecksumCRC32C:    pointer.NotBlank(cfg.Config.ChecksumCRC32C),
		ChecksumSHA1:      pointer.NotBlank(cfg.Config.ChecksumSHA1),
		ChecksumSHA256:    pointer.NotBlank(cfg.Config.ChecksumSHA256),

		ContentDisposition: pointer.NotBlank(cfg.Config.ContentDisposition),
		ContentEncoding:    pointer.NotBlank(cfg.Config.ContentEncoding),
		ContentLanguage:    pointer.NotBlank(cfg.Config.ContentLanguage),
		ContentMD5:         pointer.NotBlank(cfg.Config.ContentMD5),
		ContentType:        pointer.NotBlank(cfg.Config.ContentType),

		ExpectedBucketOwner: pointer.NotBlank(cfg.Config.ExpectedBucketOwner),
		Expires:             pointer.Time(cfg.Config.Expires),

		GrantFullControl: pointer.NotBlank(cfg.Config.GrantFullControl),
		GrantRead:        pointer.NotBlank(cfg.Config.GrantRead),
		GrantReadACP:     pointer.NotBlank(cfg.Config.GrantReadACP),
		GrantWriteACP:    pointer.NotBlank(cfg.Config.GrantWriteACP),

		Key: pointer.NotBlank(cfg.Config.Key),

		ObjectLockLegalHoldStatus: pointer.NotBlank(cfg.Config.ObjectLockLegalHoldStatus),
		ObjectLockMode:            pointer.NotBlank(cfg.Config.ObjectLockMode),
		ObjectLockRetainUntilDate: pointer.Time(cfg.Config.ObjectLockRetainUntilDate),

		RequestPayer: pointer.NotBlank(cfg.Config.RequestPayer),

		SSECustomerAlgorithm:    pointer.NotBlank(cfg.Config.SSECustomerAlgorithm),
		SSECustomerKey:          pointer.NotBlank(cfg.Config.SSECustomerKey),
		SSECustomerKeyMD5:       pointer.NotBlank(cfg.Config.SSECustomerKeyMD5),
		SSEKMSEncryptionContext: pointer.NotBlank(cfg.Config.SSEKMSEncryptionContext),
		SSEKMSKeyId:             pointer.NotBlank(cfg.Config.SSEKMSKeyId),

		ServerSideEncryption:    pointer.NotBlank(cfg.Config.ServerSideEncryption),
		StorageClass:            pointer.NotBlank(cfg.Config.StorageClass),
		Tagging:                 pointer.NotBlank(cfg.Config.Tagging),
		WebsiteRedirectLocation: pointer.NotBlank(cfg.Config.WebsiteRedirectLocation),
	}

	metadata := map[string]*string{}
	for k, v := range cfg.Config.Metadata {
		metadata[k] = &v
	}
	input.Metadata = metadata

	return input
}

// DeleteInput constructs an s3.DeleteObjectInput for deleting an object from S3.
//
// This method accepts a bucket name, object key, and optional configuration parameters.
// If no parameters are provided, default values are used.
//
// @param bucket The name of the S3 bucket to delete the object from.
// @param key The key of the object to delete.
// @param params Optional configuration parameters for the delete operation.
// @return A pointer to an s3.DeleteObjectInput with the configured values.
func DeleteInput(bucket string, key string, params ...Delete) *s3.DeleteObjectInput {
	cfg := Delete{}
	if len(params) > 0 {
		cfg = params[0]
	}

	return &s3.DeleteObjectInput{
		Bucket:                    pointer.NotBlank(bucket),
		BypassGovernanceRetention: pointer.NotFalse(cfg.BypassGovernanceRetention),
		ExpectedBucketOwner:       pointer.NotBlank(cfg.ExpectedBucketOwner),
		Key:                       pointer.NotBlank(key),
		MFA:                       pointer.NotBlank(cfg.MFA),
		RequestPayer:              pointer.NotBlank(cfg.RequestPayer),
		VersionId:                 pointer.NotBlank(cfg.Version),
	}
}

// UploadInput constructs an s3manager.UploadInput for uploading an object to S3.
//
// This method accepts a bucket name and optional configuration parameters.
// If no parameters are provided, default values are used.
//
// @param bucket The name of the S3 bucket to upload the object to.
// @param params Optional configuration parameters for the upload.
// @return A pointer to an s3manager.UploadInput with the configured values.
func UploadInput(bucket string, params ...Upload) *s3manager.UploadInput {
	cfg := Upload{}
	if len(params) > 0 {
		cfg = params[0]
	}

	input := &s3manager.UploadInput{
		Bucket:           pointer.NotBlank(bucket),
		ACL:              pointer.NotBlank(cfg.ACL),
		Body:             cfg.Body,
		BucketKeyEnabled: pointer.NotFalse(cfg.BucketKeyEnabled),
		CacheControl:     pointer.NotBlank(cfg.CacheControl),

		ChecksumAlgorithm: pointer.NotBlank(cfg.ChecksumAlgorithm),
		ChecksumCRC32:     pointer.NotBlank(cfg.ChecksumCRC32),
		ChecksumCRC32C:    pointer.NotBlank(cfg.ChecksumCRC32C),
		ChecksumSHA1:      pointer.NotBlank(cfg.ChecksumSHA1),
		ChecksumSHA256:    pointer.NotBlank(cfg.ChecksumSHA256),

		ContentDisposition: pointer.NotBlank(cfg.ContentDisposition),
		ContentEncoding:    pointer.NotBlank(cfg.ContentEncoding),
		ContentLanguage:    pointer.NotBlank(cfg.ContentLanguage),
		ContentMD5:         pointer.NotBlank(cfg.ContentMD5),
		ContentType:        pointer.NotBlank(cfg.ContentType),

		ExpectedBucketOwner: pointer.NotBlank(cfg.ExpectedBucketOwner),
		Expires:             pointer.Time(cfg.Expires),

		GrantFullControl: pointer.NotBlank(cfg.GrantFullControl),
		GrantRead:        pointer.NotBlank(cfg.GrantRead),
		GrantReadACP:     pointer.NotBlank(cfg.GrantReadACP),
		GrantWriteACP:    pointer.NotBlank(cfg.GrantWriteACP),

		Key: pointer.NotBlank(cfg.Key),

		ObjectLockLegalHoldStatus: pointer.NotBlank(cfg.ObjectLockLegalHoldStatus),
		ObjectLockMode:            pointer.NotBlank(cfg.ObjectLockMode),
		ObjectLockRetainUntilDate: pointer.Time(cfg.ObjectLockRetainUntilDate),

		RequestPayer: pointer.NotBlank(cfg.RequestPayer),

		SSECustomerAlgorithm:    pointer.NotBlank(cfg.SSECustomerAlgorithm),
		SSECustomerKey:          pointer.NotBlank(cfg.SSECustomerKey),
		SSECustomerKeyMD5:       pointer.NotBlank(cfg.SSECustomerKeyMD5),
		SSEKMSEncryptionContext: pointer.NotBlank(cfg.SSEKMSEncryptionContext),
		SSEKMSKeyId:             pointer.NotBlank(cfg.SSEKMSKeyId),

		ServerSideEncryption:    pointer.NotBlank(cfg.ServerSideEncryption),
		StorageClass:            pointer.NotBlank(cfg.StorageClass),
		Tagging:                 pointer.NotBlank(cfg.Tagging),
		WebsiteRedirectLocation: pointer.NotBlank(cfg.WebsiteRedirectLocation),
	}

	metadata := map[string]*string{}
	for k, v := range cfg.Metadata {
		metadata[k] = &v
	}
	input.Metadata = metadata

	return input
}
