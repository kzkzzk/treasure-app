package usecase

import "treasure-app/backend/domain"

type TagInteractor struct {
	TagRepository TagRepository
}

func (interactor *TagInteractor) Add(t domain.Tag) (tag domain.Tag, err error) {
	identifier, err := interactor.TagRepository.Store(t)
	if err != nil {
		return
	}
	tag, err = interactor.TagRepository.FindById(identifier)
	return
}

func (interactor *TagInteractor) Tags() (tags domain.Tags, err error) {
	tags, err = interactor.TagRepository.FindAll()
	return
}
