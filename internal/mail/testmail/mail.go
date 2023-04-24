package testmail

// Mail ...
type Mail struct{}

// New ...
func New() *Mail {
	return &Mail{}
}

// Verify ...
func (m *Mail) Verify(email, code string) error {
	return nil
}

// CommentNotify ...
func (m *Mail) CommentNotify(email, user string) error {
	return nil
}

// LikeNotify ...
func (m *Mail) LikeNotify(email, user string) error {
	return nil
}
