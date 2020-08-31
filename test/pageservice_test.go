package test

import (
	"github.com/rmpurp/knowhow/dao"
	"github.com/rmpurp/knowhow/services/impl"
	"os"
	"testing"
)

func TestEditPageService(t *testing.T) {
	_ = os.Remove("./test.db")
	connection, err := dao.ConnectAndInitialize("./test.db")
	if err != nil {
		t.Fatal(err)
	}

	title := "Mitochondria"

	//pageSet := impl.PageSet{
	//	Page:    &models.Page{},
	//	Version: &models.Version{},
	//	Content: &models.PageContent{
	//		Title:   "Mitochondria",
	//		Article: "The powerhouse of the cell",
	//	},
	//}

	editService := impl.MockEditorService{}
	pageService := impl.NewDBPageService(editService, connection)

	pageSet, err := pageService.CreatePage(title)

	if err != nil {
		t.Fatal(err)
	}

	if pageSet.Version.PageContentID != pageSet.Content.ID {
		t.Errorf("Inconsistent content IDs")
	}

	if pageSet.Version.PageID != pageSet.Page.ID {
		t.Errorf("Inconsistent page IDs")
	}

	if pageSet.Content.Title != title {
		t.Errorf("Title %v is not %v", pageSet.Content.Title, title)
	}

	if pageSet.Content.Article != "edited" {
		t.Errorf("Wrong article contents")
	}

	fetchedSets, err := pageService.FetchPagesBySearch("mitochondria")
	if err != nil {
		t.Fatal(err)
	}

	if len(fetchedSets) != 1 {
		t.Errorf("Wrong number of fetched items. Fetched items: %d", len(fetchedSets))
	}

	fetchedSet := fetchedSets[0]

	if *fetchedSet.Content != *pageSet.Content {
		t.Errorf("Fetched set not the same: %+v != %+v", fetchedSet.Content, pageSet.Content)
	}

	if *fetchedSet.Page != *pageSet.Page {
		t.Errorf("Fetched set not the same: %+v != %+v", fetchedSet.Page, pageSet.Page)
	}

	if fetchedSet.Version.IsCurrentVersion != pageSet.Version.IsCurrentVersion ||
		fetchedSet.Version.ID != pageSet.Version.ID ||
		fetchedSet.Version.PageID != pageSet.Version.PageID ||
		!fetchedSet.Version.DateCreated.Equal(pageSet.Version.DateCreated) ||
		fetchedSet.Version.PageContentID != pageSet.Version.PageContentID {
		t.Errorf("Fetched set not the same: %+v != %+v", fetchedSet.Version, pageSet.Version)
	}

	pageSet, err = pageService.EditPage(pageSet)

}
