package dao

import (
	"database/sql"
	"errors"
	"github.com/rmpurp/knowhow/models"
	"time"
)

type VersionDao interface {
	InsertOrUpdate(version *models.Version, connection *sql.Tx) error
	Delete(version *models.Version, connection *sql.Tx) error
	GetByID(id int64, connection *sql.Tx) (*models.Version, error)
}

type VersionDaoImpl struct{}

// Inserts or updates a page version in the database.
// If version is set as current version, unsets the current version flag for all others.
// Be careful if you have existing version objects as they may become out of date.
func (dao VersionDaoImpl) InsertOrUpdate(version *models.Version, connection *sql.Tx) error {
	if version.IsCurrentVersion {
		query := "UPDATE versions SET isCurrentVersion = false WHERE pageID = ?"
		_, err := connection.Exec(query, version.PageID)
		if err != nil {
			return err
		}
	}

	if version.IsInserted {
		query := "UPDATE versions SET dateCreated = ?, pageID = ?, pageContentID = ?, isCurrentVersion = ? WHERE id = ?"
		_, err := connection.Exec(
			query,
			version.DateCreated.Unix(),
			version.PageID,
			version.PageContentID,
			version.IsCurrentVersion,
			version.ID,
		)
		return err
	} else {
		query := "INSERT INTO versions (dateCreated, pageID, pageContentID, isCurrentVersion) VALUES (?, ?, ?, ?)"
		result, err := connection.Exec(
			query,
			version.DateCreated.Unix(),
			version.PageID,
			version.PageContentID,
			version.IsCurrentVersion,
		)

		if err != nil {
			return err
		}

		version.ID, err = result.LastInsertId()
		if err != nil {
			return err
		}

		version.IsInserted = true
		return nil
	}
}

func (dao VersionDaoImpl) Delete(version *models.Version, connection *sql.Tx) error {
	if !version.IsInserted {
		return errors.New("this version hasn't been inserted yet")
	}
	query := "DELETE FROM versions WHERE id = ?"
	_, err := connection.Exec(query, version.ID)
	return err
}

func (dao VersionDaoImpl) GetByID(id int64, connection *sql.Tx) (*models.Version, error) {
	query := "SELECT id, dateCreated, pageID, pageContentID, isCurrentVersion FROM versions WHERE id = ?"
	rows, err := connection.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var newID, dateCreated, pageID, pageContentID int64
		var isCurrentVersion bool
		err = rows.Scan(&newID, &dateCreated, &pageID, &pageContentID, &isCurrentVersion)

		if err != nil {
			return nil, err
		}

		version := models.Version{
			ID:               newID,
			DateCreated:      time.Unix(dateCreated, 0),
			PageID:           pageID,
			PageContentID:    pageContentID,
			IsInserted:       true,
			IsCurrentVersion: isCurrentVersion,
		}

		return &version, nil
	}

	return nil, errors.New("no object with that id")
}
