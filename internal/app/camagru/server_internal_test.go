package camagru

import (
	"bytes"
	"encoding/json"
	"github.com/bemmanue/camagru/internal/model"
	"github.com/bemmanue/camagru/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleIndex(t *testing.T) {
	s := newServer(teststore.New())

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"username": "username",
				"email":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"username": "user",
				"email":    "user@example.org",
				"password": "password",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "invalid",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sign_up", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TestUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store)

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"username": u.Username,
				"password": u.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"username": u.Username,
				"password": "invalid",
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "invalid",
			payload: map[string]string{
				"username": "invalid",
				"password": u.Password,
			},
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sign_in", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}
