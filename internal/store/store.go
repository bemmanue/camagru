package store

// Store ...
type Store interface {
	Post() PostRepository
	User() UserRepository
	Image() ImageRepository
	Like() LikeRepository
}
