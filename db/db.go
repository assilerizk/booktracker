package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
    "booktracker/config"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("postgres", config.DBUrl)
    if err != nil {
        log.Fatal("Error connecting to DB:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("DB unreachable:", err)
    }

    log.Println("Connected to database.")
}
