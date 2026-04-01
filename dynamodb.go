package main
import (
    "context"
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var DynamoClient *dynamodb.Client

func InitDynamo() {
    region := os.Getenv("AWS_REGION")
    if region == "" {
        panic("AWS_REGION is missing in env")
    }

    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(region),
    )
    if err != nil {
        panic("Failed to load AWS config: " + err.Error())
    }

    DynamoClient = dynamodb.NewFromConfig(cfg)
    fmt.Println("✅ DynamoDB client initialized for region:", region)
}

func TestDynamoConnection() {
    res, err := DynamoClient.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
    if err != nil {
        fmt.Println("❌ DynamoDB connection failed:", err)
        return
    }

    fmt.Println("✅ Connected to DynamoDB! Tables:", res.TableNames)
}