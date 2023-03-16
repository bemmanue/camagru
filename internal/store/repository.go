package store

import "github.com/bemmanue/camagru/internal/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
	FindByUsername(string) (*model.User, error)
}
