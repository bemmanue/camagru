package model

import "time"

// Post ...
type Post struct {
	ID         int
	ImageID    int
	AuthorID   int
	UploadTime time.Time
}

// PostData ...
type PostData struct {
	ID         int
	ImageID    int
	AuthorID   int
	UploadTime time.Time

	Author          string
	TimeSinceUpload string
	CommentsCount   int
	Comments        []string
	LikesCount      int
	LikeStatus      string
}
