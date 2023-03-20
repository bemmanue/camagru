package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db              *sql.DB
	userRepository  *UserRepository
	imageRepository *ImageRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Image() store.ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = &ImageRepository{
		store: s,
	}

	return s.imageRepository
}
