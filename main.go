package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "")

	router := gin.Default()

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
