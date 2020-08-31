package services

import (
	"github.com/rmpurp/knowhow/models"
)

type PageService interface {
	CreatePage(title string) (*PageSet, error)
	EditPage(pageSet PageSet) (PageSet, error)
	FetchPagesBySearch(searchTerm string) ([]*PageSet, error)
}

type PageSet struct {
	Page    *models.Page
	Version *models.Version
	Content *models.PageContent
}
