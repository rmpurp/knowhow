package dao

import (
	"database/sql"
	"errors"
	"github.com/rmpurp/knowhow/models"
)

type PageDao interface {
	InsertOrUpdate(user *models.Page, connection *sql.Tx) error
	Delete(user *models.Page, connection *sql.Tx) error
	GetByID(id int64, connection *sql.Tx) (*models.Page, error)
}

type PageDaoImpl struct{}

func (dao PageDaoImpl) InsertOrUpdate(page *models.Page, connection *sql.Tx) error {
	if page.IsInserted {
		return errors.New("page already inserted and is immutable")
	} else {
		query := "INSERT INTO pages (id) VALUES (null)"
		result, err := connection.Exec(query)
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

func (dao PageDaoImpl) Delete(user *models.Page, connection *sql.Tx) error {
	if !user.IsInserted {
		return errors.New("this page hasn't been inserted yet")
	}
	query := "DELETE FROM pages WHERE id = ?"
	_, err := connection.Exec(query, user.ID)
	return err
}

func (dao PageDaoImpl) GetByID(id int64, connection *sql.Tx) (*models.Page, error) {
	query := "SELECT id FROM pages WHERE id = ?"
	rows, err := connection.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var newID int64
		err = rows.Scan(&newID)

		if err != nil {
			return nil, err
		}

		page := models.Page{ID: newID, IsInserted: true}

		return &page, nil
	}

	return nil, errors.New("no object with that id")
}
