package access1

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"news-spider/src/utils"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", utils.SQLDB)
	if err != nil {
		panic(err)
	}
	return
}
