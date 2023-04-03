package mail

type Mail interface {
	Verify(email, code string) error
	CommentNotify(email, user string) error
	LikeNotify(email, user string) error
}
