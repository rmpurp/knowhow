package main

import (
	"database/sql"
	"github.com/gobuffalo/packr"
	_ "github.com/mattn/go-sqlite3"
)

func createFixtures(db *sql.DB) error {
	box := packr.NewBox("./database_scripts")
	script, err := box.FindString("create_fixtures.sql")

	if err != nil {
		return err
	}

	_, err = db.Exec(script)

	if err != nil {
		return err
	}

	return nil
}

func connectAndInitialize() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./debug.db")

	if err != nil {
		return nil, err
	}

	box := packr.NewBox("./database_scripts")

	script, err := box.FindString("create_database.sql")

	if err != nil {
		return database, err
	}

	_, err = database.Exec(script)

	if err != nil {
		return database, err
	}

	return database, nil
}

func addArticle(title string) bool {
	return true
}
