package camagru

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCamagru_HandleIndex(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	s.HandleIndex().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "hello")
}
