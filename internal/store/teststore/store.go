package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type Store struct {
	userRepository  *UserRepository
	imageRepository *ImageRepository
	likeRepository  *LikeRepository
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

func (s *Store) Like() store.LikeRepository {
	if s.likeRepository != nil {
		return s.likeRepository
	}

	s.likeRepository = &LikeRepository{
		store: s,
		likes: make(map[int]*model.Like),
	}

	return s.likeRepository
}
