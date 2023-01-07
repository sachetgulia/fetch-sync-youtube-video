package handlers

import (
	"fmt"
	"net/http"

	Controller "youtubesync/controllers"

	"github.com/gin-gonic/gin"
)

func SearchHandler(ctx *gin.Context) {
	var request map[string]interface{}
	err := ctx.BindJSON(&request)
	if err != nil {
		fmt.Println("error in serialising request data")
	}
	response, err := Controller.SearchController(request)
	if err != nil {
		fmt.Println("error in controller", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, response)
}

func GetDataPaginated(ctx *gin.Context) {
	var request map[string]interface{}
	err := ctx.BindJSON(&request)
	if err != nil {
		fmt.Println("error in serialising request data")
	}
	response, err := Controller.GetPaginatedDataController(request)
	if err != nil {
		fmt.Println("error in controller", err)
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, response)
}
