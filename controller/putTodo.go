package controller

import (
	"net/http"
	"todo/model"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Update todo data
func PutTodo(context *gin.Context) {

	//Initialize data todoRequest
	var data TodoRequest
	// Defining request parameter to get todo id
	ReqparamId := context.Param("idTodo")
	idTodo := cast.ToInt(ReqparamId)

	// Binding request body json to request body struct
	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return

	}

	// Initiate models todo
	todo := model.Todo{}
	// Querying find todo data by todo id from request parameter using dbfind
	todoById := db.Where(db.Where("id= ?", idTodo)).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Todo not found"})
		return
	}
	// Matching todo request with todo models

	todo.Name = data.Name
	todo.Description = data.Description
	// Update new todo data using DBsave
	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Can't update todo"})
		return
	}
	// Matching result to todo response struct
	var dataResponse TodoResponse
	dataResponse.Id = todo.Id
	dataResponse.Name = todo.Name
	dataResponse.Description = todo.Description

	// Creating http response
	context.JSON(http.StatusCreated, dataResponse)
}
