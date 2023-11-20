package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/students/:id", controllers.ShowStudent)
	r.GET("/students/cpf/:cpf", controllers.SearchStudentByCpf)
	r.POST("/students", controllers.CreateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)

	r.GET("/:name", controllers.Greetings)

	r.Run()
}
