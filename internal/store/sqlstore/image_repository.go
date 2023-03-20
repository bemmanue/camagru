package sqlstore

import "github.com/bemmanue/camagru/internal/model"

// ImageRepository ...
type ImageRepository struct {
	store *Store
}

// Create ...
func (r *ImageRepository) Create(i *model.Image) error {

	if err := r.store.db.QueryRow(
		"insert into images (user_id, filename, upload_time) values ($1, $2, $3) returning id",
		i.UserID,
		i.Filename,
		i.UploadTime,
	).Scan(&i.ID); err != nil {
		return err
	}

	return nil
}
