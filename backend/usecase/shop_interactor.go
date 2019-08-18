package usecase

import (
	"treasure-app/backend/domain"
)

type ShopInteractor struct {
	ShopRepository ShopRepository
	TagRepository  TagRepository
}

func (interactor *ShopInteractor) Add(u domain.Shop, tagIds []int) (shopDetail domain.ShopDetail, err error) {
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

	shop, err := interactor.ShopRepository.FindById(identifier)
	if err != nil {
		return
	}

	tags, err := interactor.TagRepository.FindByShopId(identifier)
	if err != nil {
		return
	}

	shopDetail = domain.ShopDetail{
		Shop: shop,
		Tags: tags,
	}
	return
}

func (interactor *ShopInteractor) Shops() (shopDetails []domain.ShopDetail, err error) {
	shops, err := interactor.ShopRepository.FindAll()
	for _, shop := range shops {
		tags, _ := interactor.TagRepository.FindByShopId(shop.ID)
		shopDetail := domain.ShopDetail{
			Shop: shop,
			Tags: tags,
		}
		shopDetails = append(shopDetails, shopDetail)
	}
	return
}

func (interactor *ShopInteractor) ShopById(identifier int) (shopDetail domain.ShopDetail, err error) {
	shop, err := interactor.ShopRepository.FindById(identifier)
	if err != nil {
		return
	}

	tags, err := interactor.TagRepository.FindByShopId(identifier)
	if err != nil {
		return
	}

	shopDetail = domain.ShopDetail{
		Shop: shop,
		Tags: tags,
	}
	return
}
