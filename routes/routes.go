package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.ShowStudent)
	r.POST("/students", controllers.CreateStudent)

	r.GET("/:name", controllers.Greetings)

	r.Run()
}
