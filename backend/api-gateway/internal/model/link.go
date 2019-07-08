package model

import "time"

type Link struct {
	ID        string
	LongLink  string
	ShortLink string
	Clicks    int
	CreatedAt time.Time
}
