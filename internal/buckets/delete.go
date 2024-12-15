package buckets

import "github.com/aws/aws-sdk-go/service/s3"

func (m *Module) Delete(bucket string) (*s3.DeleteBucketOutput, error) {
	input := &s3.DeleteBucketInput{
		Bucket: &bucket,
	}

	return m.Sdk.DeleteBucket(input)
}

func (m *Module) DeleteIfOwner(owner, bucket string) (*s3.DeleteBucketOutput, error) {
	input := &s3.DeleteBucketInput{
		ExpectedBucketOwner: &owner,
		Bucket:              &bucket,
	}

	return m.Sdk.DeleteBucket(input)
}
