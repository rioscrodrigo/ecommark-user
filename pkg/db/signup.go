package db

import (
	"ecommark-user/pkg/model"
	"ecommark-user/pkg/tools"
)

func SignUp(sign model.SignUp) error {
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	_, err = Db.Exec("INSERT INTO users (user_email, user_uuid, user_dateadd) VALUES (?, ?, ?)", sign.UserEmail, sign.UserUUID, tools.MySQLDateFormat())
	if err != nil {
		return err
	}
	return nil
}
