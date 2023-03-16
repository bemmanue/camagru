package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type UserRepository struct {
	store *Store
	users []map[string]*model.User
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[0][u.Email] = u
	r.users[1][u.Username] = u
	u.ID = len(r.users[0])

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[0][email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// FindByUsername ...
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u, ok := r.users[1][username]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
