package person

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUserService(context *gin.Context, database *gorm.DB) {
	id := context.Param("id")
	var found, err = GetUser(id, database)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, found)
}

func ListUserService(context *gin.Context, database *gorm.DB) {
	email, limitStr, offsetStr :=
		context.Query("email"),
		context.DefaultQuery("limit", "10"),
		context.DefaultQuery("offset", "0")

	var list, err = ListUser(email, limitStr, offsetStr, database)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, list)
}

func CreatePersonService(context *gin.Context, database *gorm.DB) {
	var next Person
	if err := context.BindJSON(&next); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := CreatePerson(&next, database); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	context.JSON(http.StatusOK, next)
}

func UpdatePersonService(context *gin.Context, database *gorm.DB) {
	var next Person
	if err := context.BindJSON(&next); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := UpdatePerson(&next, database); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, next)
}

func DeletePersonService(context *gin.Context, database *gorm.DB) {
	id := context.Param("id")
	var deleted, err = DeletePerson(id, database)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, deleted)
}
