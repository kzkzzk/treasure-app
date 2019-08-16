package usecase

import "treasure-app/backend/domain"

type LikeInteractor struct {
	LikeRepository LikeRepository
}

func (interactor *LikeInteractor) Add(l domain.Like) (like domain.Like, err error) {
	like, err = interactor.LikeRepository.Store(l)
	if err != nil {
		return
	}
	return
}

func (interactor *LikeInteractor) Delete(l domain.Like) (err error) {
	err = interactor.LikeRepository.Destroy(l)
	return
}
