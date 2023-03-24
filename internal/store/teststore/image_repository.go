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

// SelectImages ...
func (r *ImageRepository) SelectImages() ([]model.Image, error) {
	var images []model.Image

	for _, image := range r.images {
		images = append(images, *image)
	}

	return images, nil
}

// SelectUserImages ...
func (r *ImageRepository) SelectUserImages(userID int) ([]model.Image, error) {
	var images []model.Image

	for _, image := range r.images {
		if image.UserID == userID {
			images = append(images, *image)
		}
	}

	return images, nil
}

// SelectImagesPage ...
func (r *ImageRepository) SelectImagesPage(page int) ([]model.Image, error) {
	var images []model.Image

	for _, image := range r.images {
		if len(images) > pageSize {
			break
		}
		images = append(images, *image)
	}

	return images, nil
}
