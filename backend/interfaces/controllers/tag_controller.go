package controllers

import (
	"treasure-app/backend/domain"
	"treasure-app/backend/interfaces/database"
	"treasure-app/backend/usecase"
)

type TagController struct {
	Interactor usecase.TagInteractor
}

func NewTagController(sqlHandler database.SqlHandler) *TagController {
	return &TagController{
		Interactor: usecase.TagInteractor{
			TagRepository: &database.TagRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TagController) Create(c Context) {
	t := domain.Tag{}
	c.Bind(&t)
	Tag, err := controller.Interactor.Add(t)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, Tag)
}

func (controller *TagController) Index(c Context) {
	Tags, err := controller.Interactor.Tags()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, Tags)
}
