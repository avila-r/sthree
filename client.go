package sthree

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/internal/buckets"
	"github.com/avila-r/sthree/internal/objects"
)

type Sthree struct {
	Provider client.ConfigProvider
	Sdk      *s3.S3
	Buckets  *buckets.Module
}

func Connect(provider client.ConfigProvider, cfgs ...*aws.Config) *Sthree {
	s3 := s3.New(provider, cfgs...)

	return &Sthree{
		Provider: provider,
		Sdk:      s3,
		Buckets: &buckets.Module{
			Sdk: s3,
		},
	}
}

func New(provider client.ConfigProvider, cfgs ...*aws.Config) *Sthree {
	return Connect(provider, cfgs...)
}

func Client(provider client.ConfigProvider, cfgs ...*aws.Config) *Sthree {
	return Connect(provider, cfgs...)
}

func NewClient(provider client.ConfigProvider, cfgs ...*aws.Config) *Sthree {
	return Connect(provider, cfgs...)
}

func (c *Sthree) Bucket(name string) *objects.Module {
	return &objects.Module{
		Bucket: name,
		Sdk:    c.Buckets.Sdk,
	}
}

func (c *Sthree) In(name string) *objects.Module {
	return c.Bucket(name)
}
