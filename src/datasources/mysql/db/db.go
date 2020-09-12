package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	client *sql.DB

	username = "root"
	password = "root"
	host     = "localhost:3306"
	schema   = "oprec"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		username, password, host, schema,
	)

	var err error
	client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}

func DbConn() *sql.DB {
	return client
}
