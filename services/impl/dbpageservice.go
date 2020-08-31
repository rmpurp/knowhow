package impl

import (
	"database/sql"
	"github.com/rmpurp/knowhow/dao"
	"github.com/rmpurp/knowhow/models"
	"github.com/rmpurp/knowhow/services"
	"time"
)

type DBPageService struct {
	editorService services.EditorService
	connection    *sql.DB
}

func NewDBPageService(editorService services.EditorService, connection *sql.DB) *DBPageService {
	return &DBPageService{
		editorService: editorService,
		connection:    connection,
	}
}

func (pageService DBPageService) CreatePage(title string) (services.PageSet, error) {
	newText, err := pageService.editorService.EditText("")
	if err != nil {
		return services.PageSet{}, err
	}

	pageSet := services.PageSet{
		Page:    &models.Page{},
		Version: &models.Version{DateCreated: time.Now(), IsCurrentVersion: true},
		Content: &models.PageContent{
			Title:   title,
			Article: newText,
		},
	}

	tx, err := pageService.connection.Begin()
	if err != nil {
		return services.PageSet{}, err
	}

	pageSet, err = pageService.insertPageIntoDB(pageSet, tx)
	if err != nil {
		return services.PageSet{}, err
	}
	err = tx.Commit()

	if err != nil {
		return services.PageSet{}, err
	}

	return pageSet, nil
}

func (pageService DBPageService) EditPage(pageSet services.PageSet) (services.PageSet, error) {
	newText, err := pageService.editorService.EditText(pageSet.Content.Article)
	if err != nil {
		return services.PageSet{}, err
	}

	newContent := &models.PageContent{
		Title:   pageSet.Content.Title,
		Article: newText,
	}

	// TODO: Create Time Service
	newVersion := &models.Version{
		DateCreated:      time.Now(),
		PageID:           pageSet.Page.ID,
		PageContentID:    newContent.ID,
		IsCurrentVersion: true,
	}

	pageSet.Version = newVersion
	pageSet.Content = newContent

	tx, err := pageService.connection.Begin()
	pageSet, err = pageService.insertPageIntoDB(pageSet, tx)
	err = tx.Commit()
	if err != nil {
		return services.PageSet{}, err
	}

	return pageSet, nil
}

func (pageService DBPageService) insertPageIntoDB(pageSet services.PageSet, tx *sql.Tx) (services.PageSet, error) {
	versionDao := dao.VersionDaoImpl{}
	pageContentDao := dao.PageContentDaoImpl{}

	if !pageSet.Page.IsInserted {
		pageDao := dao.PageDaoImpl{}
		err := pageDao.InsertOrUpdate(pageSet.Page, tx)
		if err != nil {
			return services.PageSet{}, err
		}
	}

	err := pageContentDao.InsertOrUpdate(pageSet.Content, tx)
	if err != nil {
		return services.PageSet{}, err
	}

	pageSet.Version.PageID = pageSet.Page.ID
	pageSet.Version.PageContentID = pageSet.Content.ID

	err = versionDao.InsertOrUpdate(pageSet.Version, tx)

	if err != nil {
		return services.PageSet{}, err
	}

	return pageSet, nil
}

func (pageService DBPageService) FetchPagesBySearch(searchTerm string) ([]*services.PageSet, error) {
	tx, err := pageService.connection.Begin()
	if err != nil {
		return nil, err
	}
	pageContentDao := &dao.PageContentDaoImpl{}
	results, err := pageContentDao.FindBySearchString(searchTerm, tx)

	if err != nil {
		return nil, err
	}

	var pageSets []*services.PageSet

	versionDao := &dao.VersionDaoImpl{}
	pageDao := &dao.PageDaoImpl{}

	// TODO: Fix n + 1 problem
	for _, pageContent := range results {
		version, err := versionDao.GetByPageContentID(pageContent.ID, tx)
		if err != nil {
			return nil, err
		}
		page, err := pageDao.GetByID(version.PageID, tx)
		if err != nil {
			return nil, err
		}

		pageSets = append(pageSets, &services.PageSet{Page: page, Version: version, Content: pageContent})
	}

	return pageSets, nil
}
