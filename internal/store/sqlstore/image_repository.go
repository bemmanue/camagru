package sqlstore

import (
	"database/sql"
	"fmt"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
	"time"
)

const (
	pageSize = 10
)

// ImageRepository ...
type ImageRepository struct {
	Store *Store
}

// Create ...
func (r *ImageRepository) Create(i *model.Image) error {

	if err := r.Store.db.QueryRow(
		"insert into images (name, extension, path, upload_time, user_id) values ($1, $2, $3, $4, $5) returning id",
		i.Name,
		i.Extension,
		i.Path,
		i.UploadTime,
		i.UserID,
	).Scan(&i.ID); err != nil {
		return err
	}

	return nil
}

// FindByName ...
func (r *ImageRepository) FindByName(name string) (*model.Image, error) {
	img := &model.Image{}
	if err := r.Store.db.QueryRow(
		"select id, name, extension, path, upload_time, user_id from images where name = $1",
		name,
	).Scan(
		&img.ID,
		&img.Name,
		&img.Extension,
		&img.Path,
		&img.UploadTime,
		&img.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return img, nil
}

// SelectImages ...
func (r *ImageRepository) SelectImages() ([]model.Image, error) {
	query := "select images.id, name, extension, path, upload_time, images.user_id, username, count(likes.id) as like_count " +
		"from images " +
		"join users on images.user_id = users.id " +
		"left join likes on images.id = likes.image_id " +
		"group by images.id, users.username, images.upload_time " +
		"order by images.upload_time desc"

	rows, err := r.Store.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var images []model.Image
	for rows.Next() {
		image := model.Image{}
		if err := rows.Scan(
			&image.ID,
			&image.Name,
			&image.Extension,
			&image.Path,
			&image.UploadTime,
			&image.UserID,
			&image.Username,
			&image.Likes,
		); err != nil {
			return nil, err
		}

		image.TimeSinceUpload = CountTimeSinceUpload(image.UploadTime)

		images = append(images, image)
	}

	return images, nil
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

// SelectUserImages ...
func (r *ImageRepository) SelectUserImages(userID int) ([]model.Image, error) {
	query := fmt.Sprintf(
		"select images.id, name, extension, path, upload_time, images.user_id, username "+
			"from images join camagru.public.users on images.user_id = users.id "+
			"where images.user_id = %d "+
			"order by upload_time desc", userID)

	rows, err := r.Store.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var images []model.Image
	for rows.Next() {
		image := model.Image{}
		if err := rows.Scan(
			&image.ID,
			&image.Name,
			&image.Extension,
			&image.Path,
			&image.UploadTime,
			&image.UserID,
			&image.Username,
		); err != nil {
			return nil, err
		}

		image.TimeSinceUpload = CountTimeSinceUpload(image.UploadTime)

		images = append(images, image)
	}

	return images, nil
}

// SelectImagesPage ...
func (r *ImageRepository) SelectImagesPage(page int) ([]model.Image, error) {
	query := fmt.Sprintf("select images.id, name, extension, path, upload_time, images.user_id, username "+
		"from images join camagru.public.users on images.user_id = users.id "+
		"order by upload_time desc "+
		"limit %d offset %d", pageSize, pageSize*(page-1))

	rows, err := r.Store.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var images []model.Image
	for rows.Next() {
		image := model.Image{}
		if err := rows.Scan(
			&image.ID,
			&image.Name,
			&image.Extension,
			&image.Path,
			&image.UploadTime,
			&image.UserID,
			&image.Username,
		); err != nil {
			return nil, err
		}

		image.TimeSinceUpload = CountTimeSinceUpload(image.UploadTime)

		images = append(images, image)
	}

	return images, nil
}
