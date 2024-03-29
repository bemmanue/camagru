package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type UserRepository struct {
	store  *Store
	users  map[int]*model.User
	images map[int]*model.Image
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u
	return nil
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// FindByUsername ...
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// FindByUsernameVerified ...
func (r *UserRepository) FindByUsernameVerified(username string) (*model.User, error) {
	for _, user := range r.users {
		if user.Username == username && user.EmailVerified == true {
			return user, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// UsernameExists ...
func (r *UserRepository) UsernameExists(username string) (bool, error) {
	for _, user := range r.users {
		if user.Username == username {
			return true, nil
		}
	}

	return false, nil
}

// EmailExists ...
func (r *UserRepository) EmailExists(email string) (bool, error) {
	for _, user := range r.users {
		if user.Email == email {
			return true, nil
		}
	}

	return false, nil
}

// VerifyEmail ...
func (r *UserRepository) VerifyEmail(email string) error {
	for _, user := range r.users {
		if user.Email == email {
			user.EmailVerified = true
			return nil
		}
	}

	return store.ErrRecordNotFound
}

// UpdateLikeNotify ...
func (r *UserRepository) UpdateLikeNotify(id int, value bool) error {
	u, ok := r.users[id]
	if !ok {
		return store.ErrRecordNotFound
	}
	
	u.LikeNotify = value

	return nil
}

// UpdateCommentNotify ...
func (r *UserRepository) UpdateCommentNotify(id int, value bool) error {
	u, ok := r.users[id]
	if !ok {
		return store.ErrRecordNotFound
	}

	u.CommentNotify = value

	return nil
}
