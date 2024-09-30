package controller

import (
	"net/http"
	"todo/model"

	"github.com/gin-gonic/gin"
)

// Getting all todo datas
func GetTodo(context *gin.Context) {
	//Initialize new var to assign data want to get
	var datas []model.Todo
	// Querying to find todo datas.
	err := db.Find(&datas)

	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Err": "Err to getting data"})
	}
	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"Status":  "200",
		"Message": "Ok then",
		"To do ":  datas,
	})

}
