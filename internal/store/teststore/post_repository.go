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
