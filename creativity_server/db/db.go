package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pradip"
	dbname   = "UserInterestApp"
)

var db *sql.DB

func DBConnetion() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Db, err := sql.Open("postgres", psqlInfo)
	fmt.Println(Db)
	if err != nil {
		fmt.Println("Unsuccessful to connect")
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	db = Db
}

func GetDB() *sql.DB {
	return db
}
