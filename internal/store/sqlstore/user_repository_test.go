package store_test

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.New(db)
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.New(db)
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.New(db)
	username := "username"

	_, err := s.User().FindByUsername(username)
	assert.Error(t, err)

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByUsername(username)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
