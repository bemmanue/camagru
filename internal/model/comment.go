package model

import "time"

// Comment ...
type Comment struct {
	ID           int
	PostID       int
	AuthorID     int
	Author       string
	CommentText  string
	CreationTime time.Time
}
