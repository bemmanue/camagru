package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type ImageRepository struct {
	store  *Store
	images map[int]*model.Image
}

// Create ...
func (r *ImageRepository) Create(i *model.Image) error {
	i.ID = len(r.images) + 1
	r.images[i.ID] = i
	return nil
}

// Find ...
func (r *ImageRepository) Find(id int) (*model.Image, error) {
	i, ok := r.images[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return i, nil
}
