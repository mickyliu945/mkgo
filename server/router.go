package server

import (
	"github.com/gin-gonic/gin"
	"mkgo/controller"
	"mkgo/middleware/jwtauth"
	"mkgo/mkconfig"
	"net/http"
)

const rootApiName = "/api/v1"

func GetRouter() *gin.Engine {
	router := gin.Default()

	if mkconfig.Config.MLGO.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	rootRouter := router.Group(rootApiName)
	{
		user := new(controller.UserController)
		rootRouter.POST("/register", user.Register)
		rootRouter.POST("/login", user.Login)
		rootRouter.POST("/logout", user.Logout)
	}

	userGroup := router.Group(rootApiName + "/user")
	{
		userGroup.Use(jwtauth.JWTAuth())
		user := new(controller.UserController)
		userGroup.GET("/findUser", user.GetUser)
	}

	router.LoadHTMLGlob("./public/html/*")
	router.Static("/public", "./public")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name": mkconfig.Config.MLGO.Name,
		})
	})
	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	return router
}
