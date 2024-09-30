package controller

import (
	"net/http"
	"todo/connection"
	"todo/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Define DB client
var db *gorm.DB = connection.ConnectionDB()

// Todo struct for request body
type TodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Defining to do response
type TodoResponse struct {
	TodoRequest
	Id int `json:"id"`
}

// Create todo data to database by run this function
func PostTodo(context *gin.Context) {
	//Intialize data request
	var data TodoRequest

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&data); err != nil {
		//Return status and JSON string if err
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Matching todo models struct with todo request struct
	todo := model.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description
	// Querying to database (if error then bad reponse )
	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": result.Error.Error()}) // Thay đổi ở đây
	}
	// Matching result vs todo to create response
	//Call response
	var response TodoResponse

	response.Name = todo.Name
	response.Description = todo.Description

	// Creating http response with status create
	context.JSON(http.StatusCreated, response)

}
