package mail

type Mail interface {
	Verify(email, code string) error
	CommentNotify(email, commentAuthor string) error
	LikeNotify(email, likeAuthor string) error
}
