package teststore_test

import (
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImageRepository_Create(t *testing.T) {
	s := teststore.New()
	i := model.TestImage(t)
	assert.NoError(t, s.Image().Create(i))
	assert.NotNil(t, i)
}
