package store

// Store ...
type Store interface {
	Post() PostRepository
	User() UserRepository
	Image() ImageRepository
	Comment() CommentRepository
	Like() LikeRepository
	Verify() VerifyRepository
}
