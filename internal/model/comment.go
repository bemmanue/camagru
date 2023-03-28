package model

import "time"

// Comment ...
type Comment struct {
	ID         int
	AuthorID   int
	Author     string
	Text       string
	UploadTime time.Time
}