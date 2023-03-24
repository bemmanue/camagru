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

// Find ...
func (r *LikeRepository) Find(imageID, userID int) (*model.Like, error) {
	for _, like := range r.likes {
		if like.ImageID == imageID && like.UserID == userID {
			return like, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
