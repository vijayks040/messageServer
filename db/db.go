package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DbProp struct {
	Host   string
	Port   int
	User   string
	Pass   string
	DbName string
}

/*
Name: getConn
Input: Db connection properties
returns: DB connection and an error
*/
func GetConn(prop DbProp) (db *sql.DB, err error) {
	pgsql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", prop.Host, prop.Port, prop.User, prop.Pass, prop.DbName)
	db, err = sql.Open("postgres", pgsql)
	if err != nil {
		log.Println("error while connecting db", err)
	}
	return db, err
}
