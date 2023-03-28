package sqlstore

import (
	"github.com/bemmanue/camagru/internal/model"
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
