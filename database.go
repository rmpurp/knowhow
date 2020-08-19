package main

import (
    "log"
    "github.com/gobuffalo/packr"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func connectAndInitialize() *sql.DB {
    database, err := sql.Open("sqlite3", "./debug.db")

    if err != nil {
        log.Fatal(err)
    }

    box := packr.NewBox("./database_scripts")
    script, err := box.FindString("create_database.sql")
    database.Exec(script)

    if err != nil {
        log.Fatal(err)
    }

    return database
}

