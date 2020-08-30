package test

import (
	"github.com/rmpurp/knowhow/dao"
	"github.com/rmpurp/knowhow/models"
	"os"
	"testing"
	"time"
)

func TestVersionCreation(t *testing.T) {
	_ = os.Remove("./test.db")

	connection, err := dao.ConnectAndInitialize("./test.db")
	if err != nil {
		t.Errorf("could not create connection. err: %+v", err)
	}

	dao := dao.VersionDaoImpl{}
	version := &models.Version{DateCreated: time.Unix(123417239841, 0), PageID: 42, PageContentID: 21}

	err = dao.InsertOrUpdate(version, connection)

	if err != nil {
		t.Errorf("error inserting. err: %+v", err)
	}

	if !version.IsInserted {
		t.Errorf("version is not set as inserted")
	}

	if version.ID != 1 {
		t.Errorf("page id is %d when it should be %d", version.ID, 1)
	}

	returnedVersion, err := dao.GetByID(version.ID, connection)

	if err != nil {
		t.Errorf("Fetching returned an error. err: %+v", err)
	}

	if *returnedVersion != *version {
		t.Errorf("%+v != %+v", version, returnedVersion)
	}

	connection.Close()
}
