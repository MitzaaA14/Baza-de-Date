package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbName = "robby"
	dbUser = "root"
	dbPass = "root1234"
)

var DB *sql.DB

func InitDb() error {
	dbURI := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPass, dbName)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return err
	}

	DB = db

	log.Printf("DB inited with uri: %s", dbURI)

	return nil
}

func CloseDb() {
	DB.Close()
}

func GetFinalQueryToPrint(query string, params []interface{}) string {
	for _, param := range params {
		placeholder := strings.Index(query, "?")
		if placeholder != -1 {
			query = query[:placeholder] + fmt.Sprintf("'%v'", param) + query[placeholder+1:]
		}
	}

	return query
}
