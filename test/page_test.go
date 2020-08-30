package test

import (
	"github.com/rmpurp/knowhow/dao"
	"github.com/rmpurp/knowhow/models"
	"testing"
)
import (
	"os"
)


func TestPageCreation(t *testing.T) {
	_ = os.Remove("./test.db")
	connection, err := dao.ConnectAndInitialize("./test.db")
	if err != nil {
		t.Errorf("could not create connection")
	}

	dao := dao.PageDaoImpl{}
	page := &models.Page{CurrentVersion: 4}
	err = dao.InsertOrUpdate(page, connection)

	if err != nil {
		t.Errorf("error inserting")
	}

	if !page.IsInserted {
		t.Errorf("page is not set as inserted")
	}

	if page.ID != 1 {
		t.Errorf("page id is %d when it should be %d", page.ID, 1)
	}

	returnedPage, err := dao.GetByID(page.ID, connection)

	if err != nil {
		t.Errorf("Fetching returned an error")
	}

	if *returnedPage != *page {
		t.Errorf("%+v != %+v", page, returnedPage)
	}

	connection.Close()
}