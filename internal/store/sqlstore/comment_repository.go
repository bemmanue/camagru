package sqlstore

import (
	"fmt"
	"github.com/bemmanue/camagru/internal/model"
)

const (
	commentLimit = 5
)

type CommentRepository struct {
	Store *Store
}

// Create ...
func (r *CommentRepository) Create(c *model.Comment) error {
	if err := r.Store.db.QueryRow(
		"insert into comments (post_id, author_id, comment_text, creation_time) values ($1, $2, $3, $4) returning id",
		c.PostID,
		c.AuthorID,
		c.CommentText,
		c.CreationTime,
	).Scan(&c.ID); err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) GetLastComments(postID int) ([]model.Comment, error) {
	query := fmt.Sprintf("select id, author_id, author, comment_text, creation_time, post_id "+
		"from ( "+
		"select comments.id, author_id, users.username as author, comment_text, creation_time, post_id "+
		"from comments "+
		"join users on author_id = users.id "+
		"where post_id = %d "+
		"order by creation_time desc "+
		"limit %d "+
		") as comment "+
		"order by comment.creation_time asc", postID, commentLimit)

	rows, err := r.Store.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		comment := model.Comment{}
		if err := rows.Scan(
			&comment.ID,
			&comment.AuthorID,
			&comment.Author,
			&comment.CommentText,
			&comment.CreationTime,
			&comment.PostID,
		); err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

// DeleteByPostID ...
func (r *CommentRepository) DeleteByPostID(postID int) error {
	if err := r.Store.db.QueryRow(
		"delete from comments where post_id = $1",
		postID,
	); err != nil {
		return err.Err()
	}

	return nil
}
