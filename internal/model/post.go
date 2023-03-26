package model

import "time"

// Post ...
type Post struct {
	ID           int
	ImageID      int
	AuthorID     int
	CreationTime time.Time
}

// PostData ...
type PostData struct {
	ID           int
	ImageID      int
	AuthorID     int
	CreationTime time.Time

	ImagePath       string
	Author          string
	TimeSinceUpload string
	CommentsCount   int
	Comments        []string
	LikeCount       int
	LikeStatus      string
}
