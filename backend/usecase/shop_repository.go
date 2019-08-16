package usecase

import "treasure-app/backend/domain"

type ShopRepository interface {
	Store(domain.Shop) (int, error)
	FindById(int) (domain.Shop, error)
	FindAll() (domain.Shops, error)
}
