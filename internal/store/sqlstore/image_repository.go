package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

const (
	pageSize = 10
)

// ImageRepository ...
type ImageRepository struct {
	store *Store
}

// Create ...
func (r *ImageRepository) Create(i *model.Image) error {

	if err := r.store.db.QueryRow(
		"insert into images (name, extension, path, upload_time, user_id) values ($1, $2, $3, $4, $5) returning id",
		i.Name,
		i.Extension,
		i.Path,
		i.UploadTime,
		i.UserID,
	).Scan(&i.ID); err != nil {
		return err
	}

	return nil
}

// FindByName ...
func (r *ImageRepository) FindByName(name string) (*model.Image, error) {
	img := &model.Image{}
	if err := r.store.db.QueryRow(
		"select id, name, extension, path, upload_time, user_id from images where name = $1",
		name,
	).Scan(
		&img.ID,
		&img.Name,
		&img.Extension,
		&img.Path,
		&img.UploadTime,
		&img.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return img, nil
}

// SelectAllImages ...
func (r *ImageRepository) SelectAllImages() ([]string, error) {
	rows, err := r.store.db.Query("select path from images order by id")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var images []string
	image := model.Image{}

	for rows.Next() {
		rows.Scan(&image.Path)
		images = append(images, image.Path)
	}

	return images, nil
}

// GetPage ...
func (r *ImageRepository) GetPage(page int) ([]string, error) {
	rows, err := r.store.db.Query("select path from images order by id limit $1 offset $2",
		pageSize,
		pageSize*(page-1))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var images []string
	image := model.Image{}

	for rows.Next() {
		rows.Scan(&image.Path)
		images = append(images, image.Path)
	}

	return images, nil
}
