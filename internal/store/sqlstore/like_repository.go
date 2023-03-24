package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

// LikeRepository ...
type LikeRepository struct {
	Store *Store
}

func (r *LikeRepository) Create(like *model.Like) error {
	if err := r.Store.db.QueryRow(
		"insert into likes (image_id, user_id) values ($1, $2) returning id",
		like.ImageID,
		like.UserID,
	).Scan(&like.ID); err != nil {
		return err
	}

	return nil
}

func (r *LikeRepository) Delete(like *model.Like) error {
	if err := r.Store.db.QueryRow(
		"delete from likes where image_id = $1 and user_id = $2",
		like.ImageID,
		like.UserID,
	); err != nil {
		return err.Err()
	}

	return nil
}

// Find ...
func (r *LikeRepository) Find(imageID, userID int) (*model.Like, error) {
	like := &model.Like{}

	if err := r.Store.db.QueryRow(
		"select id, image_id, user_id from likes where image_id = $1 and user_id = $2",
		imageID,
		userID,
	).Scan(
		&like.ID,
		&like.ImageID,
		&like.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return like, nil
}
