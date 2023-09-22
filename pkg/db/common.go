package db

import (
	"ecommark-user/pkg/model"
	"ecommark-user/pkg/secret"
	"os"
)

var SecretModel model.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err
}
