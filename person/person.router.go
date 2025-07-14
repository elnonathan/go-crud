package person

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.RouterGroup, database *gorm.DB) {
	personRoutes := router.Group("person")

	personRoutes.GET("/:id", func(context *gin.Context) {
		GetUserService(context, database)
	})

	personRoutes.GET("/list", func(context *gin.Context) {
		ListUserService(context, database)
	})

	personRoutes.POST("/", func(context *gin.Context) {
		CreatePersonService(context, database)
	})

	personRoutes.PUT("/", func(context *gin.Context) {
		UpdatePersonService(context, database)
	})

	personRoutes.DELETE("/:id", func(context *gin.Context) {
		DeletePersonService(context, database)
	})

}
