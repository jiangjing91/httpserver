package dbaccess

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(username string, password string, dbname string) *sql.DB {
	fmt.Println(username, password, dbname)
	dbstr := username + ":" + password + "@/" + dbname
	fmt.Println(dbstr)
	conn, err := sql.Open("mysql", dbstr)

	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(100)
	conn.SetMaxIdleConns(30)
	return conn
}
