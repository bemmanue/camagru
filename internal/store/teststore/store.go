package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type Store struct {
	userRepository  *UserRepository
	imageRepository *ImageRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) Image() store.ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = &ImageRepository{
		store:  s,
		images: make(map[int]*model.Image),
	}

	return s.imageRepository
}
