package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

// PostRepository ...
type PostRepository struct {
	store *Store
	posts map[int]*model.Post
}

func (r *PostRepository) Find(id int) (*model.Post, error) {
	u, ok := r.posts[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}

// Create ...
func (r *PostRepository) Create(i *model.Post) error {
	i.ID = len(r.posts) + 1
	r.posts[i.ID] = i
	return nil
}

// GetPage ...
func (r *PostRepository) GetPage(page, userID int) ([]model.PostData, error) {
	return []model.PostData{}, nil
}

// GetUserPage ...
func (r *PostRepository) GetUserPage(page, userID int) ([]model.PostData, error) {
	return []model.PostData{}, nil
}

func (r *PostRepository) GetPageCount() (int, error) {
	return 1, nil
}

func (r *PostRepository) GetUserPageCount(userID int) (int, error) {
	return 1, nil
}

func (r *PostRepository) Delete(id int) error {
	delete(r.posts, id)

	return nil
}
