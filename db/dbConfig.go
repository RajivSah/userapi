package db

import (
	"github.com/aws/aws-sdk-go/aws"
)

type DBConfig struct {
	Region   string
	Endpoint string
}

type DBHandler struct {
	ModelName string
}

func (dbConfig DBConfig) Configuration() *aws.Config {
	config := &aws.Config{
		Region:   aws.String(dbConfig.Region),
		Endpoint: aws.String(dbConfig.Endpoint),
	}
	return config
}
