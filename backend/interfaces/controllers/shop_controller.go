package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
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
	// file upload
	file, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	dir := "uploads/" + filename
	if err := c.SaveUploadedFile(file, dir); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.Bind(&s)

	fmt.Println(s)

	createShop := domain.Shop{
		Name:    s.Name,
		Address: s.Address,
		Tel:     s.Tel,
		Image:   dir,
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
