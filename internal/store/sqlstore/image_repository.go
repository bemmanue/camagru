package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

// ImageRepository ...
type ImageRepository struct {
	Store *Store
}

// Create ...
func (r *ImageRepository) Create(i *model.Image) error {

	if err := r.Store.db.QueryRow(
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

// Find ...
func (r *ImageRepository) Find(id int) (*model.Image, error) {
	i := &model.Image{}
	if err := r.Store.db.QueryRow(
		"select id, name, extension, path, upload_time, user_id from images where id  = $1",
		id,
	).Scan(
		&i.ID,
		&i.Name,
		&i.Extension,
		&i.Path,
		&i.UploadTime,
		&i.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return i, nil
}
