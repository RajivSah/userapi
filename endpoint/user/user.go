package user

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type AWSConfig interface {
	Configuration() *aws.Config
}

type User struct {
	uid    string
	Name   string
	Age    int
	Gender string
}

func unmarshall(item map[string]*dynamodb.AttributeValue) User {
	age, _ := strconv.Atoi(*item["Age"].N)
	return User{
		uid:    *item["uid"].S,
		Name:   *item["Name"].S,
		Age:    age,
		Gender: *item["Gender"].S,
	}
}

func unmarshallList(res *dynamodb.ScanOutput) []User {
	userSlice := []User{}
	items := res.Items
	for _, user := range items {
		unmarshalledUser := unmarshall(user)
		fmt.Println(unmarshalledUser)
		userSlice = append(userSlice, unmarshalledUser)
	}
	return userSlice
}
