package store

// Store ...
type Store interface {
	User() UserRepository
	Image() ImageRepository
	Like() LikeRepository
}
