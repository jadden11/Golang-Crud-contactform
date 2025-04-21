package main

import (
	"github.com/gin-gonic/gin"
	"github.com/username/contact-form-api/controllers"
	"github.com/username/contact-form-api/database"
)

func main() {
	r := gin.Default()
	database.Connect()

	r.GET("/contacts", controllers.GetContacts)
	r.POST("/contacts", controllers.CreateContact)
	r.GET("/contacts/:id", controllers.GetContact)
	r.PUT("/contacts/:id", controllers.UpdateContact)
	r.DELETE("/contacts/:id", controllers.DeleteContact)

	r.Run() // default :8080
}
