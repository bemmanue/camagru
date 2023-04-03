package mail

type Mail interface {
	Verify(email, code string) error
	CommentNotify(email string) error
	LikeNotify(email string) error
}
