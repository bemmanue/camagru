package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type Store struct {
	userRepository *UserRepository
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
		users: make([]map[string]*model.User, 2),
	}

	s.userRepository.users[0] = make(map[string]*model.User)
	s.userRepository.users[1] = make(map[string]*model.User)

	return s.userRepository
}
