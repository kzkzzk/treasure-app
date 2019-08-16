package usecase

import (
	"treasure-app/backend/domain"
)

type ShopInteractor struct {
	ShopRepository ShopRepository
	TagRepository  TagRepository
}

func (interactor *ShopInteractor) Add(u domain.Shop, tagIds []int) (shop domain.Shop, err error) {
	// TODO: トランザクション
	identifier, err := interactor.ShopRepository.Store(u)
	if err != nil {
		return
	}

	for _, tagId := range tagIds {
		err = interactor.TagRepository.StoreShopTag(identifier, tagId)
		if err != nil {
			return
		}
	}

	shop, err = interactor.ShopRepository.FindById(identifier)
	return
}

func (interactor *ShopInteractor) Shops() (shop domain.Shops, err error) {
	shop, err = interactor.ShopRepository.FindAll()
	return
}

func (interactor *ShopInteractor) ShopById(identifier int) (shop domain.Shop, err error) {
	shop, err = interactor.ShopRepository.FindById(identifier)
	return
}
