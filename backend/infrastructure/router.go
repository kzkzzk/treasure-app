package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	gin "github.com/gin-gonic/gin"

	"treasure-app/backend/interfaces/controllers"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Authorization"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))

	router.Use(static.Serve("/uploads", static.LocalFile("./uploads", true)))

	// router.Use(authMiddleware())
	auth := NewAuth(NewSqlHandler())

	userController := controllers.NewUserController(NewSqlHandler())
	likeController := controllers.NewLikeController(NewSqlHandler())

	router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
	// いいね機能
	router.POST("users/like", auth.authMiddleware(), func(c *gin.Context) { likeController.Create(c) })
	router.DELETE("users/like", auth.authMiddleware(), func(c *gin.Context) { likeController.Delete(c) })

	shopController := controllers.NewShopController(NewSqlHandler())

	router.POST("/shops", auth.authMiddleware(), func(c *gin.Context) { shopController.Create(c) })
	router.GET("/shops", func(c *gin.Context) { shopController.Index(c) })
	router.GET("/shops/:id", func(c *gin.Context) { shopController.Show(c) })

	tagController := controllers.NewTagController(NewSqlHandler())
	// タグ機能
	router.POST("/tags", auth.authMiddleware(), func(c *gin.Context) { tagController.Create(c) })
	router.GET("/tags", func(c *gin.Context) { tagController.Index(c) })

	Router = router
}
