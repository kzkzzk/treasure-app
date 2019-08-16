package controllers

import (
	"treasure-app/backend/domain"
	"treasure-app/backend/interfaces/database"
	"treasure-app/backend/usecase"
)

type LikeController struct {
	Interactor usecase.LikeInteractor
}

func NewLikeController(sqlHandler database.SqlHandler) *LikeController {
	return &LikeController{
		Interactor: usecase.LikeInteractor{
			LikeRepository: &database.LikeRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *LikeController) Create(c Context) {
	l := domain.Like{}
	c.Bind(&l)
	Like, err := controller.Interactor.Add(l)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, Like)
}

func (controller *LikeController) Delete(c Context) {
	l := domain.Like{}
	c.Bind(&l)

	err := controller.Interactor.Delete(l)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nil)
}
