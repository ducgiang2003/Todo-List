package controller

import (
	"net/http"
	"todo/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func DeleteTodo(context *gin.Context) {
	//Intial todo model
	todo := model.Todo{}
	// Defining request parameter to get todo id
	ReqParam := context.Param("idTodo")
	idTodo := cast.ToInt(ReqParam)

	//Querrying to the db
	delete := db.Where("id = ?", idTodo).Unscoped().Delete(&todo)
	if delete.Error != nil {
		panic("Todo cant be deleted")
	}
	//Create http request (ok)
	context.JSON(http.StatusOK, gin.H{
		"Status":  "200",
		"Message": "Delete ok",
		"Data":    idTodo,
	})

}
