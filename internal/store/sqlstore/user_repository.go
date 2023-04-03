package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

// UserRepository ...
type UserRepository struct {
	Store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	if err := r.Store.db.QueryRow(
		"insert into users (username, email, encrypted_password) values ($1, $2, $3) returning id",
		u.Username,
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

// Find ...
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.Store.db.QueryRow(
		"select id, username, email, encrypted_password from users where id  = $1",
		id,
	).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil
}

// FindByUsername ...
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.Store.db.QueryRow(
		"select id, username, email, encrypted_password from users where username  = $1",
		username,
	).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil
}

// FindByUsernameVerified ...
func (r *UserRepository) FindByUsernameVerified(username string) (*model.User, error) {
	u := &model.User{}
	if err := r.Store.db.QueryRow(
		"select id, username, email, encrypted_password from users where username  = $1 and email_verified = true",
		username,
	).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.Store.db.QueryRow(
		"select id, username, email, encrypted_password from users where email  = $1",
		email,
	).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return u, nil
}

// UsernameExists ...
func (r *UserRepository) UsernameExists(username string) (bool, error) {
	var exists bool

	if err := r.Store.db.QueryRow("select "+
		"case when count(*) > 0 then true else false end as username_exists "+
		"from users where username = $1",
		username,
	).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

// EmailExists ...
func (r *UserRepository) EmailExists(email string) (bool, error) {
	var exists bool

	if err := r.Store.db.QueryRow("select "+
		"case when count(*) > 0 then true else false end as email_exists "+
		"from users where email = $1",
		email,
	).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

// VerifyEmail ...
func (r *UserRepository) VerifyEmail(email string) error {
	if err := r.Store.db.QueryRow(
		"update users set email_verified = true where email = $1", email,
	); err != nil {
		return err.Err()
	}
	return nil
}
