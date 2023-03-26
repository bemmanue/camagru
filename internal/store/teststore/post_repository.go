package teststore

import "github.com/bemmanue/camagru/internal/model"

// PostRepository ...
type PostRepository struct {
	store *Store
	posts map[int]*model.Post
}

// Create ...
func (r *PostRepository) Create(i *model.Post) error {
	i.ID = len(r.posts) + 1
	r.posts[i.ID] = i
	return nil
}

// ReadPostData ...
func (r *PostRepository) ReadPostData(userID int) ([]model.PostData, error) {
	return []model.PostData{}, nil
}
