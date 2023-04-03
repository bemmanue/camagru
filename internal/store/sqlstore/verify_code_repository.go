package sqlstore

import (
	"database/sql"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type VerifyRepository struct {
	Store *Store
}

func (r *VerifyRepository) Create(v *model.VerifyCode) error {
	if err := r.Store.db.QueryRow(
		"insert into verify_codes (code, email, user_id) values ($1, $2, $3) returning id",
		v.Code,
		v.Email,
		v.UserID,
	).Scan(&v.ID); err != nil {
		return err
	}

	return nil
}

func (r *VerifyRepository) FindByEmail(email string) (*model.VerifyCode, error) {
	v := &model.VerifyCode{}
	if err := r.Store.db.QueryRow(
		"select id, code, email, user_id from verify_codes where email  = $1",
		email,
	).Scan(
		&v.ID,
		&v.Code,
		&v.Email,
		&v.UserID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}
	return v, nil
}
