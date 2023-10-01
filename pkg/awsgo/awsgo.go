package awsgo

import (
	"context"
	"ecommark-user/cmd/logger"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func AWSInit() {
	logger.Info("Initializing AWS")
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithRegion("us-east-1"))
	if err != nil {
		logger.Error("error loading config .aws/config")
		panic("error loading config .aws/config" + err.Error())
	}
}
