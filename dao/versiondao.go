package dao

import (
	"database/sql"
	"errors"
	"github.com/rmpurp/knowhow/models"
	"time"
)

type VersionDao interface {
	InsertOrUpdate(version *models.Version, connection *sql.DB) error
	Delete(version *models.Version, connection *sql.DB) error
	GetByID(id int64, connection *sql.DB) (*models.Version, error)
}

type VersionDaoImpl struct{}

func (dao VersionDaoImpl) InsertOrUpdate(version *models.Version, connection *sql.DB) error {
	if version.IsInserted {
		query := "UPDATE versions SET dateCreated = ?, pageID = ?, pageContentID = ? WHERE id = ?"
		_, err := connection.Exec(query, version.DateCreated.Unix(), version.PageID, version.PageContentID, version.ID)
		return err
	} else {
		query := "INSERT INTO versions (dateCreated, pageID, pageContentID) VALUES (?, ?, ?)"
		result, err := connection.Exec(query, version.DateCreated.Unix(), version.PageID, version.PageContentID)

		if err != nil {
			return err
		} else {
			version.ID, err = result.LastInsertId()
			if err != nil {
				return err
			}

			version.IsInserted = true
			return nil
		}
	}
}

func (dao VersionDaoImpl) Delete(version *models.Version, connection *sql.DB) error {
	if !version.IsInserted {
		return errors.New("this version hasn't been inserted yet")
	}
	query := "DELETE FROM versions WHERE id = ?"
	_, err := connection.Exec(query, version.ID)
	return err
}

func (dao VersionDaoImpl) GetByID(id int64, connection *sql.DB) (*models.Version, error) {
	query := "SELECT id, dateCreated, pageID, pageContentID FROM versions WHERE id = ?"
	rows, err := connection.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var newID, dateCreated, pageID, pageContentID int64
		err = rows.Scan(&newID, &dateCreated, &pageID, &pageContentID)

		if err != nil {
			return nil, err
		}

		version := models.Version{
			ID:            newID,
			DateCreated:   time.Unix(dateCreated, 0),
			PageID:        pageID,
			PageContentID: pageContentID,
			IsInserted:    true,
		}

		return &version, nil
	}

	return nil, errors.New("no object with that id")
}
