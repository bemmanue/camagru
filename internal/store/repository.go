package store

import "github.com/bemmanue/camagru/internal/model"

// PostRepository ...
type PostRepository interface {
	Create(*model.Post) error
	ReadPostData(userID int) ([]model.PostData, error)
}

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindByUsername(string) (*model.User, error)
}

// ImageRepository ...
type ImageRepository interface {
	Create(*model.Image) error
}

// CommentRepository ...
type CommentRepository interface {
	Create(*model.Comment) error
}

// LikeRepository ...
type LikeRepository interface {
	Create(*model.Like) error
	Delete(*model.Like) error
	Find(imageID, userID int) (*model.Like, error)
}
