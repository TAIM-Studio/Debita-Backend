package controller

import (
	"net/http"

	"github.com/TAIM-Studio/Debita-Backend/common/datamodel"
	"github.com/TAIM-Studio/Debita-Backend/common/dbInterface"
	"github.com/gin-gonic/gin"
)

func CreateAbo() gin.HandlerFunc {
	handler := func(context *gin.Context) {
		var newAbo datamodel.Abo
		if err := context.BindJSON(&newAbo); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Corrupted Data send",
			})
			return
		}
		abo, err := dbInterface.CreateAbo(&newAbo)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		context.IndentedJSON(http.StatusCreated, abo)

	}
	return gin.HandlerFunc(handler)
}

func GetAllAbos() gin.HandlerFunc {
	handler := func(context *gin.Context) {
		abos, err := dbInterface.GetAllAbos()
		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		}
		context.IndentedJSON(http.StatusOK, abos)
	}
	return gin.HandlerFunc(handler)
}
