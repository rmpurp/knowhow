package test

import (
	"database/sql"
	"github.com/rmpurp/knowhow/dao"
	"github.com/rmpurp/knowhow/models"
	"os"
	"testing"
	"time"
)

func confirmVersion(t *testing.T, versionID int64, isCurrentVersion bool, tx *sql.Tx) {
	t.Run("confirm", func(t *testing.T) {
		versionDao := dao.VersionDaoImpl{}

		versionFetched, err := versionDao.GetByID(versionID, tx)
		if err != nil {
			t.Errorf("Error when fetching version with ID=%v", versionID)
		} else if versionFetched.IsCurrentVersion != isCurrentVersion {
			t.Errorf("version %+v current version status should be %v", versionFetched, isCurrentVersion)
		}
	})
}

func TestVersionWithPage(t *testing.T) {
	_ = os.Remove("./test.db")

	connection, err := dao.ConnectAndInitialize("./test.db")
	if err != nil {
		t.Errorf("could not create connection. err: %+v", err)
	}

	tx, err := connection.Begin()

	pageDao := dao.PageDaoImpl{}
	page1 := &models.Page{}
	page2 := &models.Page{}
	pageDao.InsertOrUpdate(page1, tx)
	pageDao.InsertOrUpdate(page2, tx)

	tx.Commit()

	versionDao := dao.VersionDaoImpl{}

	version1page1 := &models.Version{
		DateCreated:      time.Unix(1000, 0),
		PageID:           page1.ID,
		PageContentID:    1,
		IsCurrentVersion: true,
	}

	version2page1 := &models.Version{
		DateCreated:      time.Unix(2000, 0),
		PageID:           page1.ID,
		PageContentID:    2,
		IsCurrentVersion: true,
	}

	version1page2 := &models.Version{
		DateCreated:      time.Unix(3000, 0),
		PageID:           page2.ID,
		PageContentID:    3,
		IsCurrentVersion: true,
	}

	version2page2 := &models.Version{
		DateCreated:      time.Unix(4000, 0),
		PageID:           page2.ID,
		PageContentID:    4,
		IsCurrentVersion: true,
	}

	tx, err = connection.Begin()

	err = versionDao.InsertOrUpdate(version1page1, tx)
	if err != nil {
		t.Errorf("error when inserting %+v; err: %v", version2page2, err)
	}

	err = versionDao.InsertOrUpdate(version2page1, tx)
	if err != nil {
		t.Errorf("error when inserting %+v; err: %v", version2page2, err)
	}

	err = versionDao.InsertOrUpdate(version1page2, tx)
	if err != nil {
		t.Errorf("error when inserting %+v; err: %v", version2page2, err)
	}

	tx.Commit()

	tx, err = connection.Begin()
	confirmVersion(t, version1page1.ID, false, tx)
	confirmVersion(t, version2page1.ID, true, tx)
	confirmVersion(t, version1page2.ID, true, tx)
	tx.Commit()

	tx, err = connection.Begin()
	err = versionDao.InsertOrUpdate(version2page2, tx)
	if err != nil {
		t.Errorf("error when inserting %+v; err: %v", version2page2, err)
	}
	tx.Commit()

	tx, err = connection.Begin()
	confirmVersion(t, version1page1.ID, false, tx)
	confirmVersion(t, version2page1.ID, true, tx)
	confirmVersion(t, version1page2.ID, false, tx)
	confirmVersion(t, version2page2.ID, true, tx)
	tx.Commit()
}

func TestVersionCreation(t *testing.T) {
	_ = os.Remove("./test.db")

	connection, err := dao.ConnectAndInitialize("./test.db")
	if err != nil {
		t.Errorf("could not create connection. err: %+v", err)
	}

	dao := dao.VersionDaoImpl{}
	version := &models.Version{
		DateCreated:      time.Unix(123417239841, 0),
		PageID:           42,
		PageContentID:    21,
		IsCurrentVersion: true,
	}

	tx, err := connection.Begin()
	err = dao.InsertOrUpdate(version, tx)
	tx.Commit()

	if err != nil {
		t.Errorf("error inserting. err: %+v", err)
	}

	if !version.IsInserted {
		t.Errorf("version is not set as inserted")
	}

	if version.ID != 1 {
		t.Errorf("page id is %d when it should be %d", version.ID, 1)
	}

	tx, err = connection.Begin()
	returnedVersion, err := dao.GetByID(version.ID, tx)
	tx.Commit()

	if err != nil {
		t.Errorf("Fetching returned an error. err: %+v", err)
	}

	if *returnedVersion != *version {
		t.Errorf("%+v != %+v", version, returnedVersion)
	}

	connection.Close()
}
