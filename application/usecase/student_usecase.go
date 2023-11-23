package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/domain/student/entity"
)

type StudentUseCase struct {
	studentRepository entity.StudentRepository
}

func NewWithDependencies(studentRepository entity.StudentRepository) *StudentUseCase {
	return &StudentUseCase{studentRepository: studentRepository}
}

func (uc *StudentUseCase) FindAll() ([]entity.Student, error) {
	return uc.studentRepository.FindAll()
}

func (uc *StudentUseCase) FindById(ctx *gin.Context) (entity.Student, error) {
	id := ctx.Param("id")

	if id == "" {
		return entity.Student{}, nil
	}

	return uc.studentRepository.FindById(id)
}

func (uc *StudentUseCase) Create(ctx *gin.Context) (entity.Student, error) {
	var student entity.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		return student, err
	}

	if err := entity.ValidateStudent(&student); err != nil {
		return student, err
	}

	return uc.studentRepository.Create(student)
}

func (uc *StudentUseCase) Update(ctx *gin.Context) (entity.Student, error) {
	var student entity.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		return student, err
	}

	if err := entity.ValidateStudent(&student); err != nil {
		return student, err
	}

	return uc.studentRepository.Update(student)
}

func (uc *StudentUseCase) Delete(ctx *gin.Context) error {
	id := ctx.Param("id")

	if id == "" {
		return nil
	}

	return uc.studentRepository.Delete(id)
}

func (uc *StudentUseCase) SearchStudentByCpf(ctx *gin.Context) (entity.Student, error) {
	cpf := ctx.Param("cpf")
	return uc.studentRepository.SearchStudentByCpf(cpf)
}
