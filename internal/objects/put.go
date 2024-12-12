package objects

import (
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/pointer"
)

type Put struct {
	Config ObjectDetails
}

func (m *Module) Put(body interface{}, params ...Put) (*s3.PutObjectOutput, error) {
	input := &s3.PutObjectInput{
		Bucket: &m.Bucket,
		Body:   ToReadSeeker(body),
	}
	if len(params) > 0 {
		input = fromPutParams(m.Bucket, body, params[0])
	}

	return m.Sdk.PutObject(input)
}

func fromPutParams(bucket string, body interface{}, params Put) *s3.PutObjectInput {
	input := &s3.PutObjectInput{
		Bucket:           &bucket,
		ACL:              pointer.NotBlank(params.Config.ACL),
		Body:             ToReadSeeker(body),
		BucketKeyEnabled: &params.Config.BucketKeyEnabled,
		CacheControl:     pointer.NotBlank(params.Config.CacheControl),

		ChecksumAlgorithm: pointer.NotBlank(params.Config.ChecksumAlgorithm),
		ChecksumCRC32:     pointer.NotBlank(params.Config.ChecksumCRC32),
		ChecksumCRC32C:    pointer.NotBlank(params.Config.ChecksumCRC32C),
		ChecksumSHA1:      pointer.NotBlank(params.Config.ChecksumSHA1),
		ChecksumSHA256:    pointer.NotBlank(params.Config.ChecksumSHA256),

		ContentDisposition: pointer.NotBlank(params.Config.ContentDisposition),
		ContentEncoding:    pointer.NotBlank(params.Config.ContentEncoding),
		ContentLanguage:    pointer.NotBlank(params.Config.ContentLanguage),
		ContentMD5:         pointer.NotBlank(params.Config.ContentMD5),
		ContentType:        pointer.NotBlank(params.Config.ContentType),

		ExpectedBucketOwner: pointer.NotBlank(params.Config.ExpectedBucketOwner),
		Expires:             pointer.Time(params.Config.Expires),

		GrantFullControl: pointer.NotBlank(params.Config.GrantFullControl),
		GrantRead:        pointer.NotBlank(params.Config.GrantRead),
		GrantReadACP:     pointer.NotBlank(params.Config.GrantReadACP),
		GrantWriteACP:    pointer.NotBlank(params.Config.GrantWriteACP),

		Key: pointer.NotBlank(params.Config.Key),

		ObjectLockLegalHoldStatus: pointer.NotBlank(params.Config.ObjectLockLegalHoldStatus),
		ObjectLockMode:            pointer.NotBlank(params.Config.ObjectLockMode),
		ObjectLockRetainUntilDate: pointer.Time(params.Config.ObjectLockRetainUntilDate),

		RequestPayer: pointer.NotBlank(params.Config.RequestPayer),

		SSECustomerAlgorithm:    pointer.NotBlank(params.Config.SSECustomerAlgorithm),
		SSECustomerKey:          pointer.NotBlank(params.Config.SSECustomerKey),
		SSECustomerKeyMD5:       pointer.NotBlank(params.Config.SSECustomerKeyMD5),
		SSEKMSEncryptionContext: pointer.NotBlank(params.Config.SSEKMSEncryptionContext),
		SSEKMSKeyId:             pointer.NotBlank(params.Config.SSEKMSKeyId),

		ServerSideEncryption:    pointer.NotBlank(params.Config.ServerSideEncryption),
		StorageClass:            pointer.NotBlank(params.Config.StorageClass),
		Tagging:                 pointer.NotBlank(params.Config.Tagging),
		WebsiteRedirectLocation: pointer.NotBlank(params.Config.WebsiteRedirectLocation),
	}

	metadata := map[string]*string{}
	for k, v := range params.Config.Metadata {
		metadata[k] = &v
	}
	input.Metadata = metadata

	return input
}
