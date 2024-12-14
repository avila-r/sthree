package sthree

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/avila-r/sthree/internal/buckets"
	"github.com/avila-r/sthree/internal/objects"
	"github.com/avila-r/sthree/internal/requests"
)

type Sthree struct {
	Provider client.ConfigProvider
	Sdk      *s3.S3
	Buckets  *buckets.Module
	Requests *requests.Module
}

func Connect(provider client.ConfigProvider, cfgs ...Config) *Sthree {
	s3 := s3.New(provider, unwrapConfig(cfgs...)...)

	return &Sthree{
		Provider: provider,
		Sdk:      s3,
		Buckets: &buckets.Module{
			Sdk: s3,
		},
		Requests: &requests.Module{
			Sdk: s3,
		},
	}
}

func Session(cfgs ...Config) (*Sthree, error) {
	sess, err := session.NewSession(unwrapConfig(cfgs...)...)

	if err != nil {
		return nil, err
	}

	return Connect(sess), nil
}

func New(provider client.ConfigProvider, cfgs ...Config) *Sthree {
	return Connect(provider, cfgs...)
}

func Client(provider client.ConfigProvider, cfgs ...Config) *Sthree {
	return Connect(provider, cfgs...)
}

func NewClient(provider client.ConfigProvider, cfgs ...Config) *Sthree {
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

func unwrapConfig(c ...Config) []*aws.Config {
	if len(c) > 0 {
		return []*aws.Config{
			c[0].ToAWSConfig(),
		}
	}

	// Empty
	return []*aws.Config{}
}
