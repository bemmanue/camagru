package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type Store struct {
	postRepository    *PostRepository
	userRepository    *UserRepository
	imageRepository   *ImageRepository
	commentRepository *CommentRepository
	likeRepository    *LikeRepository
	verifyRepository  *VerifyRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) Post() store.PostRepository {
	if s.postRepository != nil {
		return s.postRepository
	}

	s.postRepository = &PostRepository{
		store: s,
		posts: make(map[int]*model.Post),
	}

	return s.postRepository
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

func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}

	s.commentRepository = &CommentRepository{
		store:    s,
		comments: make(map[int]*model.Comment),
	}

	return s.commentRepository
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

func (s *Store) Verify() store.VerifyRepository {
	if s.verifyRepository != nil {
		return s.verifyRepository
	}

	s.verifyRepository = &VerifyRepository{
		store: s,
		codes: make(map[int]*model.VerifyCode),
	}

	return s.verifyRepository
}
