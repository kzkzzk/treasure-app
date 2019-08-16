package controllers

import (
	"strconv"
	"treasure-app/backend/domain"
	"treasure-app/backend/interfaces/database"
	"treasure-app/backend/usecase"
)

type ShopController struct {
	Interactor usecase.ShopInteractor
}

func NewShopController(sqlHandler database.SqlHandler) *ShopController {
	return &ShopController{
		Interactor: usecase.ShopInteractor{
			ShopRepository: &database.ShopRepository{
				SqlHandler: sqlHandler,
			},
			TagRepository: &database.TagRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ShopController) Create(c Context) {
	s := domain.RequestCreateShop{}
	c.Bind(&s)

	createShop := domain.Shop{
		Name:    s.Name,
		Address: s.Address,
		Tel:     s.Tel,
		Image:   s.Image,
	}

	shop, err := controller.Interactor.Add(createShop, s.TagIDs)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, shop)
}

func (controller *ShopController) Index(c Context) {
	shops, err := controller.Interactor.Shops()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, shops)
}

func (controller *ShopController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	shop, err := controller.Interactor.ShopById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, shop)
}
