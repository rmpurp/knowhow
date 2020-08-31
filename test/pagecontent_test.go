package test

import (
	"github.com/rmpurp/knowhow/dao"
	"github.com/rmpurp/knowhow/models"
	"log"
	"os"
	"testing"
)

func TestPageContentSearch(t *testing.T) {
	_ = os.Remove("./test.db")

	connection, err := dao.ConnectAndInitialize("./test.db")

	if err != nil {
		log.Fatal(err)
	}

	pageContentDao := &dao.PageContentDaoImpl{}

	pageContent1 := &models.PageContent{
		Title:   "git subtrees",
		Article: "git subtrees are a useful feature involving repositories in repositories",
	}

	pageContent2 := &models.PageContent{
		Title:   "version control",
		Article: "version control includes things like git.",
	}
	tx, err := connection.Begin()
	err = pageContentDao.InsertOrUpdate(pageContent1, tx)
	if err != nil {
		t.Fatal(err)
	}
	err = pageContentDao.InsertOrUpdate(pageContent2, tx)
	if err != nil {
		t.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	tx, err = connection.Begin()
	matchingPageContents, err := pageContentDao.FindBySearchString("git", tx)
	if err != nil {
		t.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		t.Fatal(err)
	}

	if len(matchingPageContents) != 2 {
		t.Fatal("should have matched both pages")
	}

	if *matchingPageContents[0] != *pageContent1 {
		t.Errorf("Actual best match %+v != expected best match %+v", matchingPageContents[0], pageContent1)
	}

	if *matchingPageContents[1] != *pageContent2 {
		t.Errorf("Actual second-best match %+v != expected second-best match %+v", matchingPageContents[0], pageContent1)
	}

	tx, err = connection.Begin()
	matchingPageContents, err = pageContentDao.FindBySearchString("control", tx)
	if err != nil {
		t.Fatal(err)
	}
	err = tx.Commit()

	if len(matchingPageContents) != 1 {
		t.Fatal("should have matched one page")
	}

	if *matchingPageContents[0] != *pageContent2 {
		t.Errorf("Actual best match %+v != expected best match %+v", matchingPageContents[0], pageContent1)
	}
}

func TestPageContentCreation(t *testing.T) {
	_ = os.Remove("./test.db")

	connection, err := dao.ConnectAndInitialize("./test.db")
	if err != nil {
		log.Fatal(err)
	}

	pageDao := &dao.PageContentDaoImpl{}

	tx, err := connection.Begin()
	if err != nil {
		log.Fatal(err)
	}

	pageContent := &models.PageContent{
		Title:   "Git Subtrees",
		Article: "Git subtrees are quite useful",
	}

	err = pageDao.InsertOrUpdate(pageContent, tx)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	if pageContent.IsInserted != true {
		t.Errorf("page content is not marked as inserted")
	}

	if pageContent.ID != 1 {
		t.Errorf("page content's id is not set correctly")
	}

	tx, err = connection.Begin()
	fetchedPageContent, err := pageDao.GetByID(pageContent.ID, tx)

	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}

	if *fetchedPageContent != *pageContent {
		t.Errorf("%+v != %+v", fetchedPageContent, pageContent)
	}
}
