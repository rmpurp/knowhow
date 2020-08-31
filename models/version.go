package models

import "time"

type Version struct {
	ID               int64
	DateCreated      time.Time
	PageID           int64
	PageContentID    int64
	IsCurrentVersion bool
	IsInserted       bool
}
