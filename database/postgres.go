package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DB_CONNECTION"))
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
