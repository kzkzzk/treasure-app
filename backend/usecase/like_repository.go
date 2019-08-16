package usecase

import "treasure-app/backend/domain"

type LikeRepository interface {
	Store(domain.Like) (domain.Like, error)
	Destroy(domain.Like) error
}
