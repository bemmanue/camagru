package sqlstore

import (
	"fmt"
	"github.com/bemmanue/camagru/internal/model"
	"time"
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

		if post.Comments, err = r.Store.Comment().GetLastComments(post.ID); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// ReadUserPostData ...
func (r *PostRepository) ReadUserPostData(userID int) ([]model.PostData, error) {
	query := fmt.Sprintf(
		"select posts.id, posts.image_id, posts.author_id, posts.creation_time, "+
			"images.path, users.username, "+
			"count(likes.id) as likes_count, "+
			"(case when sum(case when likes.user_id = %d then 1 else 0 END) > 0 then 'like' else 'dislike' end) as user_like "+
			"from posts join images on posts.image_id = images.id "+
			"join users on posts.author_id = users.id "+
			"left join likes on images.id = likes.image_id "+
			"where posts.author_id = %d "+
			"group by posts.id, users.username, images.path "+
			"order by posts.creation_time desc", userID, userID)

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

		if post.Comments, err = r.Store.Comment().GetLastComments(post.ID); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// CountTimeSinceUpload ...
func CountTimeSinceUpload(uploadTime time.Time) string {
	var result string

	timeNow := time.Now()
	timeSpan := timeNow.Sub(uploadTime)

	if timeSpan < time.Minute {
		result = fmt.Sprintf("%d seconds", int(timeSpan.Round(time.Second).Seconds()))
	} else if timeSpan < time.Hour {
		result = fmt.Sprintf("%d minutes", int(timeSpan.Round(time.Minute).Minutes()))
	} else if timeSpan < time.Hour*24 {
		result = fmt.Sprintf("%d hours", int(timeSpan.Round(time.Hour).Hours()))
	} else if timeSpan < time.Hour*24*7 {
		result = fmt.Sprintf("%d days", int(timeSpan.Round(time.Hour).Hours()/24))
	} else if timeSpan < time.Hour*24*365 {
		result = fmt.Sprintf("%d weeks", int(timeSpan.Round(time.Hour).Hours()/24/7))
	} else {
		result = fmt.Sprintf("%d years", int(timeSpan.Round(time.Hour).Hours()/24/365))
	}
	return result
}
