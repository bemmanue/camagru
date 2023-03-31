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
		"insert into likes (post_id, user_id) values ($1, $2) returning id",
		like.PostID,
		like.UserID,
	).Scan(&like.ID); err != nil {
		return err
	}

	return nil
}

func (r *LikeRepository) Delete(like *model.Like) error {
	if err := r.Store.db.QueryRow(
		"delete from likes where post_id = $1 and user_id = $2",
		like.PostID,
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
		"select id, post_id, user_id from likes where post_id = $1 and user_id = $2",
		imageID,
		userID,
	).Scan(
		&like.ID,
		&like.PostID,
		&like.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return like, nil
}
