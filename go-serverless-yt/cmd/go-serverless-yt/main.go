package main

import (
	"os"

	"github.com/alvintoh/go-programming-projects/go-serverless-yt/internal/app/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)

const tableName = "go-serverless-yt"

func handlerRoutes(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handler.GetUser(req, tableName, dynaClient)
	case "POST":
		return handler.CreateUser(req, tableName, dynaClient)
	case "PUT":
		return handler.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handler.DeleteUser(req, tableName, dynaClient)
	default:
		return handler.UnhandledMethod()
	}
}

func main() {
	//setupConfig()
	//startServer(r)
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}
	dynaClient = dynamodb.New(awsSession)
	lambda.Start(handlerRoutes)
}
