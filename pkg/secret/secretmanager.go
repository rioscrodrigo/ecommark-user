package secret

import (
	"ecommark-user/cmd/logger"
	"ecommark-user/pkg/awsgo"
	"ecommark-user/pkg/model"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (model.SecretRDSJson, error) {
	var secretData model.SecretRDSJson
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		logger.Error("error getting secret")
		return secretData, err
	}
	json.Unmarshal([]byte(*key.SecretString), &secretData)
	return secretData, nil
}
