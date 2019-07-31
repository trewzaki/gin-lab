package main

import (
	"github.com/gin-gonic/gin"
	"github.com/trewzaki/gin-lab/controllers"
	"github.com/trewzaki/gin-lab/models"
)

func main() {
	models.InitialDatabase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	r.POST("/user/create", controllers.CreateUser)

	r.Run(":8080")
}
