package test

import (
	"github.com/rmpurp/knowhow/dao"
	"github.com/rmpurp/knowhow/models"
	"os"
	"testing"
)

func TestPageCreation(t *testing.T) {
	_ = os.Remove("./test.db")
	connection, err := dao.ConnectAndInitialize("./test.db")
	if err != nil {
		t.Errorf("could not create connection")
	}

	dao := dao.PageDaoImpl{}
	page := &models.Page{}

	tx, err := connection.Begin()

	err = dao.InsertOrUpdate(page, tx)

	tx.Commit()

	if err != nil {
		t.Errorf("error inserting")
	}

	if !page.IsInserted {
		t.Errorf("page is not set as inserted")
	}

	if page.ID != 1 {
		t.Errorf("page id is %d when it should be %d", page.ID, 1)
	}

	tx, err = connection.Begin()

	returnedPage, err := dao.GetByID(page.ID, tx)

	tx.Commit()

	if err != nil {
		t.Errorf("Fetching returned an error")
	}

	if returnedPage == nil {
		t.Errorf("returned page is nil")
	} else if *returnedPage != *page {
		t.Errorf("%+v != %+v", page, returnedPage)
	}

	connection.Close()
}
