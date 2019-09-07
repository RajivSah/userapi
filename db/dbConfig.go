package db

import (
	"github.com/aws/aws-sdk-go/aws"
)

type DBHandler struct {
	ModelName string
	Region    string
	Endpoint  string
}

func (dbConfig DBHandler) Configuration() *aws.Config {
	config := &aws.Config{
		Region:   aws.String(dbConfig.Region),
		Endpoint: aws.String(dbConfig.Endpoint),
	}
	return config
}
