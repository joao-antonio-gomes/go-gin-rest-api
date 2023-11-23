package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/infrastructure/rest"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("api/students", rest.ShowAllStudents)
	r.GET("api/students/:id", rest.ShowStudent)
	r.GET("api/students/cpf/:cpf", rest.SearchStudentByCpf)
	r.POST("api/students", rest.CreateStudent)
	r.DELETE("api/students/:id", rest.DeleteStudent)
	r.PATCH("api/students/:id", rest.EditStudent)

	r.Run()
}
