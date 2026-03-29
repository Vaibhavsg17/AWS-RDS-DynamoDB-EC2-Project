package main

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
)

var DynamoClient *dynamodb.Client
var TableName = "tasks"

func InitDynamo() {
cfg, err := config.LoadDefaultConfig(context.TODO(),
	config.WithRegion(os.Getenv("AWS_REGION")),

)
	if err != nil {
		log.Fatal(err)
	}

	DynamoClient = dynamodb.NewFromConfig(cfg)
	log.Println("Connected to DynamoDB ✅")
}