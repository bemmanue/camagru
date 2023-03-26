package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db              *sql.DB
	postRepository  *PostRepository
	userRepository  *UserRepository
	imageRepository *ImageRepository
	likeRepository  *LikeRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Post() store.PostRepository {
	if s.postRepository != nil {
		return s.postRepository
	}

	s.postRepository = &PostRepository{
		Store: s,
	}

	return s.postRepository
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		Store: s,
	}

	return s.userRepository
}

func (s *Store) Image() store.ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = &ImageRepository{
		Store: s,
	}

	return s.imageRepository
}

func (s *Store) Like() store.LikeRepository {
	if s.likeRepository != nil {
		return s.likeRepository
	}

	s.likeRepository = &LikeRepository{
		Store: s,
	}

	return s.likeRepository
}
