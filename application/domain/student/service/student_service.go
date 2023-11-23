package service

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/domain/student/entity"
)

type StudentService interface {
	FindAll() ([]entity.Student, error)
	FindById(ctx *gin.Context) (entity.Student, error)
	Create(ctx *gin.Context) (entity.Student, error)
	Update(student *entity.Student) (*entity.Student, error)
	Delete(id string) error
	SearchStudentByCpf(cpf string) (entity.Student, error)
}
