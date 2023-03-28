package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
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
