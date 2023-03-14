package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Username: "username",
		Email:    "user@example.org",
		Password: "password",
	}
}
