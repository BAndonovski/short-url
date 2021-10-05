package model

import "time"

type ShortLink struct {
	Short     string
	Original  string
	CreatedOn time.Time
	LastVisit time.Time
	Visits    int64
}
