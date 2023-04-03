package teststore

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
)

type VerifyRepository struct {
	store *Store
	codes map[int]*model.VerifyCode
}

func (r *VerifyRepository) Create(v *model.VerifyCode) error {
	v.ID = len(r.codes) + 1
	r.codes[v.ID] = v
	return nil
}

func (r *VerifyRepository) FindByEmail(email string) (*model.VerifyCode, error) {
	for _, code := range r.codes {
		if code.Email == email {
			return code, nil
		}
	}

	return nil, store.ErrRecordNotFound
}
