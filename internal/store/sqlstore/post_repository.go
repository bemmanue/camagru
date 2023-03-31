package sqlstore

import (
	"fmt"
	"github.com/bemmanue/camagru/internal/model"
	"time"
)

const (
	pageSize = 10
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

func (r *PostRepository) GetPageCount() (int, error) {
	var count int
	if err := r.Store.db.QueryRow(
		"select count(*) from posts",
	).Scan(&count); err != nil {
		return 0, err
	}
	return count/pageSize + 1, nil
}

func (r *PostRepository) GetUserPageCount(userID int) (int, error) {
	var count int
	if err := r.Store.db.QueryRow(
		"select count(*) from posts where author_id = $1", userID,
	).Scan(&count); err != nil {
		return 0, err
	}
	return count/pageSize + 1, nil
}

func (r *PostRepository) GetPage(page, userID int) ([]model.PostData, error) {
	query := fmt.Sprintf(
		"select posts.id, posts.image_id, posts.author_id, posts.creation_time, "+
			"images.path, users.username, "+
			"count(distinct likes.id) as likes_count, "+
			"count(distinct comments.id) as comments_count, "+
			"(case when sum(case when likes.user_id = %d then 1 else 0 END) > 0 then 'like' else 'dislike' end) as user_like "+
			"from posts join images on posts.image_id = images.id "+
			"join users on posts.author_id = users.id "+
			"left join likes on posts.id = likes.post_id "+
			"left join comments on posts.id = comments.post_id "+
			"group by posts.id, users.username, images.path, posts.creation_time "+
			"order by posts.creation_time desc "+
			"limit %d offset %d", userID, pageSize, (page-1)*pageSize)

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
			&post.CommentCount,
			&post.LikeStatus,
		); err != nil {
			return nil, err
		}

		post.TimeSinceUpload = CountTimeSinceUpload(post.CreationTime)

		// Add post comments
		if post.Comments, err = r.Store.Comment().GetLastComments(post.ID); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// GetUserPage ...
func (r *PostRepository) GetUserPage(page, userID int) ([]model.PostData, error) {
	query := fmt.Sprintf(
		"select posts.id, posts.image_id, posts.author_id, posts.creation_time, "+
			"images.path, users.username, "+
			"count(distinct likes.id) as likes_count, "+
			"count(distinct comments.id) as comment_count, "+
			"(case when sum(case when likes.user_id = %d then 1 else 0 END) > 0 then 'like' else 'dislike' end) as user_like "+
			"from posts join images on posts.image_id = images.id "+
			"join users on posts.author_id = users.id "+
			"left join likes on posts.id = likes.post_id "+
			"left join comments on posts.id = comments.post_id "+
			"where posts.author_id = %d "+
			"group by posts.id, users.username, images.path, posts.creation_time "+
			"order by posts.creation_time desc "+
			"limit %d offset %d", userID, userID, pageSize, (page-1)*pageSize)

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
			&post.CommentCount,
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
