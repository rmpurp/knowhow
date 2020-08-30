package dao

import (
	"database/sql"
	"github.com/gobuffalo/packr"
	_ "github.com/mattn/go-sqlite3"
)

func CreateFixtures(db *sql.DB) error {
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

func ConnectAndInitialize(filename string) (*sql.DB, error) {
	database, err := sql.Open("sqlite3", filename)

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
