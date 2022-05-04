package models

import "time"

type Post struct {
	Title     string
	Author    string
	Content   string
	Slug      string
	CreatedAt time.Time
}
