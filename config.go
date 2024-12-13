package sthree

import (
	"github.com/aws/aws-sdk-go/aws"

	"github.com/avila-r/sthree/pkg/pointer"
)

type Config struct {
	Region   string
	Advanced aws.Config
}

func (c Config) ToAWSConfig() *aws.Config {
	cfg := c.Advanced
	cfg.Region = pointer.NotBlank(c.Region)
	return &c.Advanced
}
