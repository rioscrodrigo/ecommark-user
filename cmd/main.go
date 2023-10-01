package main

import (
	"context"
	"ecommark-user/cmd/logger"
	"ecommark-user/pkg/awsgo"
	"ecommark-user/pkg/db"
	"ecommark-user/pkg/model"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	logger.Info("Starting lambda...")
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	logger.Info("Lambda execution started")
	awsgo.AWSInit()

	if !ValidateParams() {
		logger.Error("missing parameters")
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
		logger.Error("error reading secret")
		return event, err
	}
	err = db.SignUp(datos)
	return event, err
}

func ValidateParams() bool {
	var hasParameter bool
	_, hasParameter = os.LookupEnv("SecretName")
	return hasParameter
}
