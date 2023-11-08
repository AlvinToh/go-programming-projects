package main

import (
	"fmt"

	"github.com/alvintoh/go-programming-projects/lambda-yt-example/internal/app/model"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent = model.MyEvent
type MyResponse = model.MyResponse

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
