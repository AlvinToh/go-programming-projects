package handler

import (
	"net/http"

	"github.com/alvintoh/go-programming-projects/go-serverless-yt/internal/app/model"
	"github.com/alvintoh/go-programming-projects/go-serverless-yt/internal/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := model.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return util.ApiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return util.ApiResponse(http.StatusOK, result)
	}

	result, err := model.FetchUsers(tableName, dynaClient)
	if err != nil {
		return util.ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return util.ApiResponse(http.StatusOK, result)

}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := model.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return util.ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return util.ApiResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := model.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return util.ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return util.ApiResponse(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	err := model.DeleteUser(req, tableName, dynaClient)

	if err != nil {
		return util.ApiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return util.ApiResponse(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return util.ApiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
