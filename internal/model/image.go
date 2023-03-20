package model

import "time"

// Image ...
type Image struct {
	ID         int
	UserID     int
	Filename   string
	UploadTime time.Time
}
