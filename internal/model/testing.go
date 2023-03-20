package model

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) *User {
	return &User{
		Username: "username",
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestImage(t *testing.T) *Image {
	return &Image{
		UserID:     1,
		UploadTime: time.Now(),
	}
}
