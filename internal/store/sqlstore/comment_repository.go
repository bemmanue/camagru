package sqlstore

import "github.com/bemmanue/camagru/internal/model"

type CommentRepository struct {
	Store *Store
}

// Create ...
func (r *CommentRepository) Create(c *model.Comment) error {
	if err := r.Store.db.QueryRow(
		"insert into comments (author_id, text, creation_time) values ($1, $2, $3) returning id",
		c.AuthorID,
		c.Text,
		c.CreationTime,
	).Scan(&c.ID); err != nil {
		return err
	}

	return nil
}
