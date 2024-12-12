package objects

import (
	"io"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/avila-r/sthree/pointer"
)

type Upload struct {
	// The readable body payload to send to S3.
	Body io.Reader
	ObjectDetails
}

func (m *Module) Upload(params ...Upload) (*s3manager.UploadOutput, error) {
	input := &s3manager.UploadInput{
		Bucket: &m.Bucket,
	}
	if len(params) > 0 {
		input = fromUploadParams(m.Bucket, params[0])
	}

	return m.Uploader.Upload(input)
}

func fromUploadParams(name string, params Upload) *s3manager.UploadInput {
	input := &s3manager.UploadInput{
		Bucket:           &name,
		ACL:              pointer.NotBlank(params.ACL),
		Body:             params.Body,
		BucketKeyEnabled: &params.BucketKeyEnabled,
		CacheControl:     pointer.NotBlank(params.CacheControl),

		ChecksumAlgorithm: pointer.NotBlank(params.ChecksumAlgorithm),
		ChecksumCRC32:     pointer.NotBlank(params.ChecksumCRC32),
		ChecksumCRC32C:    pointer.NotBlank(params.ChecksumCRC32C),
		ChecksumSHA1:      pointer.NotBlank(params.ChecksumSHA1),
		ChecksumSHA256:    pointer.NotBlank(params.ChecksumSHA256),

		ContentDisposition: pointer.NotBlank(params.ContentDisposition),
		ContentEncoding:    pointer.NotBlank(params.ContentEncoding),
		ContentLanguage:    pointer.NotBlank(params.ContentLanguage),
		ContentMD5:         pointer.NotBlank(params.ContentMD5),
		ContentType:        pointer.NotBlank(params.ContentType),

		ExpectedBucketOwner: pointer.NotBlank(params.ExpectedBucketOwner),
		Expires:             pointer.Time(params.Expires),

		GrantFullControl: pointer.NotBlank(params.GrantFullControl),
		GrantRead:        pointer.NotBlank(params.GrantRead),
		GrantReadACP:     pointer.NotBlank(params.GrantReadACP),
		GrantWriteACP:    pointer.NotBlank(params.GrantWriteACP),

		Key: pointer.NotBlank(params.Key),

		ObjectLockLegalHoldStatus: pointer.NotBlank(params.ObjectLockLegalHoldStatus),
		ObjectLockMode:            pointer.NotBlank(params.ObjectLockMode),
		ObjectLockRetainUntilDate: pointer.Time(params.ObjectLockRetainUntilDate),

		RequestPayer: pointer.NotBlank(params.RequestPayer),

		SSECustomerAlgorithm:    pointer.NotBlank(params.SSECustomerAlgorithm),
		SSECustomerKey:          pointer.NotBlank(params.SSECustomerKey),
		SSECustomerKeyMD5:       pointer.NotBlank(params.SSECustomerKeyMD5),
		SSEKMSEncryptionContext: pointer.NotBlank(params.SSEKMSEncryptionContext),
		SSEKMSKeyId:             pointer.NotBlank(params.SSEKMSKeyId),

		ServerSideEncryption:    pointer.NotBlank(params.ServerSideEncryption),
		StorageClass:            pointer.NotBlank(params.StorageClass),
		Tagging:                 pointer.NotBlank(params.Tagging),
		WebsiteRedirectLocation: pointer.NotBlank(params.WebsiteRedirectLocation),
	}

	metadata := map[string]*string{}
	for k, v := range params.Metadata {
		metadata[k] = &v
	}
	input.Metadata = metadata

	return input
}
