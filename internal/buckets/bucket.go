package buckets

import (
	"github.com/avila-r/sthree/internal/bucket"
	"github.com/avila-r/sthree/internal/objects"
)

// Bucket returns an instance of the `bucket.Module` for the specified bucket.
//
// This method is used to interact with bucket-specific operations,
// providing an interface to access configurations related to the bucket.
//
// @param name The name of the bucket to associate with the returned module.
// @return An instance of `bucket.Module` configured with the bucket name and the associated SDK.
func (m *Module) Bucket(name string) *bucket.Module {
	return &bucket.Module{
		Bucket: name,
		Sdk:    m.Sdk,
	}
}

// From returns an instance of the `objects.Module` for the specified bucket.
//
// This method is used for operations involving objects within a bucket,
// such as uploading, downloading, listing, and deleting objects.
//
// @param name The name of the bucket to associate with the returned module.
// @return An instance of `objects.Module` configured with the bucket name and the associated SDK.
func (m *Module) From(name string) *objects.Module {
	return &objects.Module{
		Bucket: name,
		Sdk:    m.Sdk,
	}
}
