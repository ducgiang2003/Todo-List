package routes

import (
	"todo/controller"

	"github.com/gin-gonic/gin"
)

//Use gin frameword to config RESTful API endpoint

func Routes() {
	routes := gin.Default()
	routes.POST("/todo", controller.PostTodo)
	routes.GET("/todo", controller.GetTodo)
	routes.PUT("/todo/:idTodo", controller.PutTodo)
	routes.DELETE("/todo/:idTodo", controller.DeleteTodo)
	//Run route
	routes.Run(":8825")
}
