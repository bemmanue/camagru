package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

const (
	pageSize = 10
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

// FindByName ...
func (r *ImageRepository) FindByName(name string) (*model.Image, error) {
	for _, image := range r.images {
		if image.Name == name {
			return image, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// SelectAllImages ...
func (r *ImageRepository) SelectAllImages() ([]string, error) {
	var images []string

	for _, image := range r.images {
		images = append(images, image.Path)
	}

	return images, nil
}

// GetPage ...
func (r *ImageRepository) GetPage(page int) ([]string, error) {
	var images []string

	for _, image := range r.images {
		if len(images) > pageSize {
			break
		}
		images = append(images, image.Path)
	}

	return images, nil
}
