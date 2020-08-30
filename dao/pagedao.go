package dao

import (
	"database/sql"
	"errors"
	"github.com/rmpurp/knowhow/models"
)

type PageDao interface {
	InsertOrUpdate(user *models.Page, connection *sql.DB) error
	Delete(user *models.Page, connection *sql.DB) error
	GetByID(id int64, connection *sql.DB) (*models.Page, error)
}

type PageDaoImpl struct { }

func (dao PageDaoImpl) InsertOrUpdate(page *models.Page, connection *sql.DB) error {
	if page.IsInserted {
		query := "UPDATE pages SET currentVersion = ? WHERE id = ?"
		_, err := connection.Exec(query, page.CurrentVersion, page.ID)
		return err
	} else {
		query := "INSERT INTO pages (currentVersion) VALUES (?)"
		result, err := connection.Exec(query, page.CurrentVersion)
		if err != nil {
			return err
		} else {
			page.ID, err = result.LastInsertId()
			if err != nil {
				return err
			}

			page.IsInserted = true
			return nil
		}
	}
}

func (dao PageDaoImpl) Delete(user *models.Page, connection *sql.DB) error {
	if !user.IsInserted {
		return errors.New("this page hasn't been inserted yet")
	}
	query := "DELETE FROM pages WHERE id = ?"
	_, err := connection.Exec(query, user.ID)
	return err
}

func (dao PageDaoImpl) GetByID(id int64, connection *sql.DB) (*models.Page, error) {
	query := "SELECT id, currentVersion FROM pages WHERE id = ?"
	rows, err := connection.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var newID int64
		var currentVersion int64
		err = rows.Scan(&newID, &currentVersion)

		if err != nil {
			return nil, err
		}

		page := models.Page{ID: newID, CurrentVersion: currentVersion, IsInserted: true}

		return &page, nil
	}

	return nil, errors.New("no object with that id")
}
