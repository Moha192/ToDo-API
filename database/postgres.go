package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable host=db_todo port=5432")
	time.Sleep(time.Second * 1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (userid SERIAL PRIMARY KEY, email VARCHAR(255) UNIQUE NOT NULL, password TEXT NOT NULL)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS tasks (taskid SERIAL PRIMARY KEY, userid INT REFERENCES users(userid), task TEXT, status BOOLEAN DEFAULT FALSE);")
	if err != nil {
		log.Fatal(err)
	}
}
