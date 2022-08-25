package routes

import (
	"todo/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRouters(router *gin.Engine) {
	//todo routers
	router.GET("/todo", controllers.GetAllTodo)
	router.POST("/todo", controllers.AddOneTodo)
	router.DELETE("/todo/:id", controllers.DeleteOneTodo)

	// user routers
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUserHandler)
}
