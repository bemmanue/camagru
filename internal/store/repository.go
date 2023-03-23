package store

import "github.com/bemmanue/camagru/internal/model"

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
	FindByName(name string) (*model.Image, error)
	SelectImages() ([]model.Image, error)
	SelectUserImages(userID int) ([]model.Image, error)
	SelectImagesPage(page int) ([]model.Image, error)
}
