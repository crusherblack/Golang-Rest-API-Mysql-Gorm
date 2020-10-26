package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"your.import/path/controllers"
	"your.import/path/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	log.Println("Starting development server at http://127.0.0.1:8000/")
	log.Println("Quit the server with CONTROL-C.")

	models.ConnectDatabase()

	r.GET("/bookings", controllers.FindAll)
	r.GET("/booking/:id", controllers.Find)
	r.POST("/booking", controllers.Create)
	r.PATCH("/booking/:id", controllers.Update)
	r.DELETE("/booking/:id", controllers.Delete)

	r.Run()
}
