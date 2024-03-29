package store

import "github.com/bemmanue/camagru/internal/model"

// PostRepository ...
type PostRepository interface {
	Find(int) (*model.Post, error)
	Delete(int) error
	Create(*model.Post) error
	GetPage(page, userID int) ([]model.PostData, error)
	GetUserPage(page, userID int) ([]model.PostData, error)
	GetPageCount() (int, error)
	GetUserPageCount(userID int) (int, error)
}

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	FindByUsername(string) (*model.User, error)
	FindByUsernameVerified(string) (*model.User, error)
	UsernameExists(string) (bool, error)
	EmailExists(string) (bool, error)
	VerifyEmail(string) error
	UpdateLikeNotify(id int, value bool) error
	UpdateCommentNotify(id int, value bool) error
}

// ImageRepository ...
type ImageRepository interface {
	Create(*model.Image) error
	Find(int) (*model.Image, error)
}

// CommentRepository ...
type CommentRepository interface {
	Create(*model.Comment) error
	GetLastComments(postID int) ([]model.Comment, error)
	DeleteByPostID(int) error
}

// LikeRepository ...
type LikeRepository interface {
	Create(*model.Like) error
	Delete(*model.Like) error
	Find(imageID, userID int) (*model.Like, error)
	DeleteByPostID(int) error
}

// VerifyRepository ...
type VerifyRepository interface {
	Create(*model.VerifyCode) error
	FindByEmail(string) (*model.VerifyCode, error)
}
