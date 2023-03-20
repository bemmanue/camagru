package sqlstore_test

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	i := model.TestImage(t)
	assert.NoError(t, s.Image().Create(i))
	assert.NotNil(t, i)
}
