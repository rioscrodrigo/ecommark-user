package tools

import "time"

func MySQLDateFormat() string {
	t := time.Now()
	return t.Format("2006-01-02T15:04:05")
}
