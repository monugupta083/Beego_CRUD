package routers

import (
	"myproject/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	//userController := new(controllers.UserController)
	var userService controllers.UserInterface
	userService = &controllers.UserController{}
	r.GET("/users", userService.Index)
	r.GET("/users/:id", userService.Show)
	r.POST("/users", userService.Create)
	r.PUT("/users/:id", userService.Update)
	r.DELETE("/users/:id", userService.Destroy)

	return r
}
