package main

import (
	"github.com/gin-gonic/gin"
	"go-crud/db"
	"go-crud/person"
)

func SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-DNS-Prefetch-Control", "off")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer")
		c.Writer.Header().Set("Permissions-Policy", "geolocation=(), microphone=()")
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		c.Next()
	}
}

func main() {
	var database, err = db.OpenDatabase()
	if err != nil {
		panic(err.Error())
		return
	}

	router := gin.Default()
	api := router.Group("/")
	router.Use(SecureHeaders())

	person.RegisterRoutes(api, database)

	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err.Error())
		return
	}

	err = router.Run(":8080")
	if err != nil {
		panic(err.Error())
		return
	}
}
