package model

// VerifyCode ...
type VerifyCode struct {
	ID     int
	Email  string
	Code   int
	UserID int
}
