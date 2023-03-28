package teststore

import "github.com/bemmanue/camagru/internal/model"

// CommentRepository ...
type CommentRepository struct {
	store    *Store
	comments map[int]*model.Comment
}

// Create ...
func (r *CommentRepository) Create(c *model.Comment) error {
	c.ID = len(r.comments) + 1
	r.comments[c.ID] = c
	return nil
}
