package usecase

import "treasure-app/backend/domain"

type TagRepository interface {
	Store(domain.Tag) (int, error)
	FindById(int) (domain.Tag, error)
	FindAll() (domain.Tags, error)
}
