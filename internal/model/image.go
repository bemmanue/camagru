package model

import "time"

// Image ...
type Image struct {
	ID              int
	Name            string
	Extension       string
	Path            string
	UploadTime      time.Time
	TimeSinceUpload string
	UserID          int
	Username        string
}
