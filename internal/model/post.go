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
	CommentCount    int
	Comments        []Comment
	LikeCount       int
	LikeStatus      string
}
