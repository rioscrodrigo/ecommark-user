package db

import (
	"database/sql"
	"ecommark-user/cmd/logger"
	"ecommark-user/pkg/model"
	"ecommark-user/pkg/secret"
	"os"
)

var SecretModel model.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		logger.Error("Error al conectar a la base de datos")
		return err
	}

	err = Db.Ping()
	if err != nil {
		logger.Error("Error al hacer ping a la base de datos")
		return err
	}
	return nil
}

func ConnStr(clave model.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = clave.Username
	authToken = clave.Password
	dbEndpoint = clave.Host
	dbName = "ecommark"
	dns := dbUser + ":" + authToken + "@tcp(" + dbEndpoint + ":3306)/" + dbName + "?allowCleartextPasswords=true"
	return dns
}
