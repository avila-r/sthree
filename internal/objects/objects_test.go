package objects_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/avila-r/sthree"
	"github.com/avila-r/sthree/pkg/mock"
)

var client = func() *sthree.Sthree {
	sess, err := session.NewSession()

	if err != nil {
		panic(err.Error())
	}

	return sthree.New(sess)
}()

func Test_ObjectOperations(t *testing.T) {
	bucket := mock.RandomBucketName()
	if _, err := client.Buckets.New(bucket); err != nil {
		t.Errorf("failed to create bucket - %v", err.Error())
	}

	t.Cleanup(func() {
		output, err := client.In(bucket).List()
		if err != nil {
			t.Errorf("failed to retrieve objects in bucket to delete them - %v", err.Error())
		}

		for _, object := range output.Contents {
			if _, err := client.In(bucket).Delete(*object.Key); err != nil {
				t.Errorf("failed to delete object - %v", err.Error())
			}
		}

		if _, err := client.Buckets.Delete(bucket); err != nil {
			t.Errorf("failed to delete bucket - %v", err.Error())
		}
	})

	id := mock.RandomObjectID()
	data := struct {
		Name string `json:"name"`
	}{"test-data"}

	_, err := client.
		In(bucket).
		Put(id, data)

	if err != nil {
		t.Errorf("failed to put object - %v", err.Error())
	}
}
