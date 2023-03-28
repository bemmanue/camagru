package sqlstore

import (
	"fmt"
	"github.com/bemmanue/camagru/internal/model"
)

// PostRepository ...
type PostRepository struct {
	Store *Store
}

// Create ...
func (r *PostRepository) Create(p *model.Post) error {

	if err := r.Store.db.QueryRow(
		"insert into posts (image_id, author_id, creation_time) values ($1, $2, $3) returning id",
		p.ImageID,
		p.AuthorID,
		p.CreationTime,
	).Scan(&p.ID); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) ReadPostData(userID int) ([]model.PostData, error) {
	query := fmt.Sprintf(
		"select posts.id, posts.image_id, posts.author_id, posts.creation_time, "+
			"images.path, users.username, "+
			"count(likes.id) as likes_count, "+
			"(case when sum(case when likes.user_id = %d then 1 else 0 END) > 0 then 'like' else 'dislike' end) as user_like "+
			"from posts join images on posts.image_id = images.id "+
			"join users on posts.author_id = users.id "+
			"left join likes on images.id = likes.image_id "+
			"group by posts.id, users.username, images.path "+
			"order by posts.creation_time desc", userID)

	rows, err := r.Store.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []model.PostData
	for rows.Next() {
		post := model.PostData{}
		if err := rows.Scan(
			&post.ID,
			&post.ImageID,
			&post.AuthorID,
			&post.CreationTime,
			&post.ImagePath,
			&post.Author,
			&post.LikeCount,
			&post.LikeStatus,
		); err != nil {
			return nil, err
		}

		post.TimeSinceUpload = CountTimeSinceUpload(post.CreationTime)

		posts = append(posts, post)
	}

	return posts, nil
}