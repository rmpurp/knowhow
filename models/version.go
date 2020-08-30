package models

import "time"

type Version struct {
	ID            int
	DateCreated   time.Time
	PageID        int
	PageContentID int
	IsInserted    bool
}
