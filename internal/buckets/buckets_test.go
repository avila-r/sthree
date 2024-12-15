package buckets_test

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/avila-r/sthree"
)

var client = func() *sthree.Sthree {
	sess, err := session.NewSession()

	if err != nil {
		panic(err.Error())
	}

	return sthree.New(sess)
}()

func Test_Create_Delete(t *testing.T) {
	// Hold created buckets
	// to allow cleaning
	created := []string{}

	t.Cleanup(func() {
		for _, bucket := range created {
			if _, err := client.Buckets.Delete(bucket); err != nil {
				t.Errorf("failed to delete bucket - %v", err.Error())
			}
		}
	})

	t.Run("default usecase", func(t *testing.T) {
		name := RandomBucketName()

		_, err := client.Buckets.Create(name)
		if err != nil {
			t.Errorf("failed to create a bucket - %v", err.Error())
		}
		created = append(created, name)
	})
}

func Test_List(t *testing.T) {
	// Hold created buckets
	// to allow cleaning
	created := []string{}

	t.Cleanup(func() {
		for _, bucket := range created {
			if _, err := client.Buckets.Delete(bucket); err != nil {
				t.Errorf("failed to delete bucket - %v", err.Error())
			}
		}
	})

	quantity := 6
	for i := 1; i < quantity; i++ {
		name := RandomBucketName()

		_, err := client.Buckets.Create(name)
		if err != nil {
			t.Errorf("failed to create bucket - %v", err.Error())
		}
		created = append(created, name)
	}

	output, err := client.Buckets.List()
	if err != nil {
		t.Errorf("failed to list buckets - %v", err.Error())
	}

	for i, bucket := range output.Buckets {
		t.Logf("[%v] Created Bucket: %v", i+1, *bucket.Name)
	}
}

func RandomBucketName() string {
	words := []string{
		"apple", "banana", "cherry", "date", "elderberry", "fig", "grape",
		"honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "papaya",
		"quince", "raspberry", "strawberry", "tangerine", "ugli", "vanilla",
		"watermelon", "xigua", "yam", "zucchini",
	}

	passphrase := []string{}
	size := 6
	for i := 0; i < size; i++ {
		passphrase = append(passphrase, words[rand.Intn(len(words))])
	}

	return strings.Join(passphrase, "-")
}
