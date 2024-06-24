package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(192.168.118.26:3306)/Tables_in_video_server")
	if err != nil {
		panic(err.Error())
	}
}
