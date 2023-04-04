package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

// LikeRepository ...
type LikeRepository struct {
	store *Store
	likes map[int]*model.Like
}

// Create ...
func (r *LikeRepository) Create(l *model.Like) error {
	l.ID = len(r.likes) + 1
	r.likes[l.ID] = l
	return nil
}

// Delete ...
func (r *LikeRepository) Delete(l *model.Like) error {
	for i, like := range r.likes {
		if like.PostID == l.PostID &&
			like.UserID == l.UserID {
			delete(r.likes, i)
		}
	}
	return nil
}

// Find ...
func (r *LikeRepository) Find(imageID, userID int) (*model.Like, error) {
	for _, like := range r.likes {
		if like.PostID == imageID && like.UserID == userID {
			return like, nil
		}
	}
	return nil, store.ErrRecordNotFound
}

// DeleteByPostID ...
func (r *LikeRepository) DeleteByPostID(postID int) error {
	for id, like := range r.likes {
		if like.PostID == postID {
			delete(r.likes, id)
		}
	}
	return nil
}
