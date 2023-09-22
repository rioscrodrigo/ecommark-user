package main

import (
	"context"
	"ecommark-user/pkg/awsgo"
	"ecommark-user/pkg/db"
	"ecommark-user/pkg/model"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AWSInit()
	if !ValidateParams() {
		fmt.Println("missing parameters")
		err := errors.New("missing parameters")
		return event, err
	}

	var datos model.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
		case "sub":
			datos.UserUUID = att
		}
	}
	err := db.ReadSecret()
	if err != nil {
		fmt.Println("error reading secret")
		return event, err
	}
	return event, nil
}

func ValidateParams() bool {
	var hasParameter bool
	_, hasParameter = os.LookupEnv("SecretName")
	return hasParameter
}
