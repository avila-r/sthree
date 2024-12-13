package sthree_test

import (
	"testing"

	"github.com/avila-r/sthree"
)

func Test_Session(t *testing.T) {
	client, err := sthree.Session()

	if err != nil {
		t.Errorf("failed to create new session - %v", err.Error())
	}

	if _, err := client.Buckets.List(); err != nil {
		t.Errorf("failed to list buckets - %v", err.Error())
	}

	all, _ := client.Buckets.List()

	t.Logf("%v existent buckets:", len(all.Buckets))

	for i, bucket := range all.Buckets {
		t.Logf("[%v] Name: %v", i+1, *bucket.Name)
	}
}
