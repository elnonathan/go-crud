package test

import (
	"github.com/gin-gonic/gin"
	"go-crud/person"
	"gorm.io/gorm"
)

func SetupTestRouter(database *gorm.DB) *gin.Engine {
	router := gin.Default()
	api := router.Group("/")
	person.RegisterRoutes(api, database)
	return router
}
