package dao

import (
	"database/sql"
	"errors"
	"github.com/rmpurp/knowhow/models"
)

type PageContentDao interface {
	InsertOrUpdate(user *models.PageContent, connection *sql.Tx) error
	Delete(pageContent *models.PageContent, connection *sql.Tx) error
	GetByID(id int64, connection *sql.Tx) (*models.PageContent, error)
	FindBySearchString(searchString string, connection *sql.Tx) (*models.PageContent, error)
}

type PageContentDaoImpl struct{}

func (dao PageContentDaoImpl) InsertOrUpdate(pageContent *models.PageContent, connection *sql.Tx) error {
	if pageContent.IsInserted {
		query := "UPDATE pageContent SET title = ?, article = ? WHERE rowid = ?"
		_, err := connection.Exec(query, pageContent.Title, pageContent.Article)
		return err
	} else {
		query := "INSERT INTO pageContent (title, article) VALUES (?, ?)"
		result, err := connection.Exec(query, pageContent.Title, pageContent.Article)
		if err != nil {
			return err
		} else {
			pageContent.ID, err = result.LastInsertId()
			if err != nil {
				return err
			}

			pageContent.IsInserted = true
			return nil
		}
	}
}

func (dao PageContentDaoImpl) Delete(pageContent *models.PageContent, connection *sql.Tx) error {
	if !pageContent.IsInserted {
		return errors.New("this PageContent hasn't been inserted yet")
	}

	query := "DELETE FROM pageContent WHERE rowid = ?"
	_, err := connection.Exec(query, pageContent.ID)
	return err
}

func (dao PageContentDaoImpl) GetByID(id int64, connection *sql.Tx) (*models.PageContent, error) {
	query := "SELECT rowid, title, article FROM pageContent WHERE rowid = ?"
	rows, err := connection.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var newID int64
		var title, article string
		err = rows.Scan(&newID, &title, &article)

		if err != nil {
			return nil, err
		}

		page := models.PageContent{ID: newID, Title: title, Article: article, IsInserted: true}

		return &page, nil
	}

	return nil, errors.New("no object with that id")
}

func (dao PageContentDaoImpl) FindBySearchString(searchString string, connection *sql.Tx) ([]*models.PageContent, error) {
	query := "SELECT c.rowid, c.title, c.article FROM pageContent c LEFT OUTER JOIN Versions v ON c.ROWID == v.pageContentID WHERE c.pageContent = ? AND COALESCE(v.isCurrentVersion, 1) order by rank"
	rows, err := connection.Query(query, searchString)

	if err != nil {
		return nil, err
	}

	var foundPageContents []*models.PageContent

	var id int64
	var title, article string

	for rows.Next() {
		err = rows.Scan(&id, &title, &article)
		if err != nil {
			return nil, err
		}

		newArticle := &models.PageContent{
			ID:         id,
			Title:      title,
			Article:    article,
			IsInserted: true,
		}

		foundPageContents = append(foundPageContents, newArticle)
	}

	return foundPageContents, nil
}
