package infrastructure

import (
	gin "github.com/gin-gonic/gin"

	"treasure-app/backend/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	shopController := controllers.NewShopController(NewSqlHandler())

	router.POST("/shops", func(c *gin.Context) { shopController.Create(c) })
	router.GET("/shops", func(c *gin.Context) { shopController.Index(c) })
	router.GET("/shops/:id", func(c *gin.Context) { shopController.Show(c) })

	Router = router
}
