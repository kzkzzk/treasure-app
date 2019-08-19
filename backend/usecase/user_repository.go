package usecase

import "treasure-app/backend/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
	Update(domain.User) error
	FindByFirebaseId(string) (domain.User, error)
}
