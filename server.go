package main

import (
	"server/controllers"
	"server/middlewares"
	"server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.JSONMiddleware())
	r.Use(middlewares.CORSMiddleware())

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	models.ConnectDatabase()

	r.Run(":4444")
}
