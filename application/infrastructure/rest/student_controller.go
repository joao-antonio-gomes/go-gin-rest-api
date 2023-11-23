package rest

import (
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/domain/student/entity"
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joao-antonio-gomes/go-gin-rest-api/database"
)

type StudentController struct {
	studentUseCase *usecase.StudentUseCase
}

var singletonStudentController *StudentController

func NewStudentController(studentUseCase *usecase.StudentUseCase) *StudentController {
	singletonStudentController = &StudentController{studentUseCase: studentUseCase}
	return singletonStudentController
}

func ShowAllStudents(ctx *gin.Context) {
	all, _ := singletonStudentController.studentUseCase.FindAll()
	ctx.JSON(http.StatusOK, all)
}

func ShowStudent(ctx *gin.Context) {
	var student entity.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	ctx.JSON(http.StatusOK, &student)
}

func CreateStudent(ctx *gin.Context) {
	var student entity.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := entity.ValidateStudent(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	ctx.JSON(http.StatusCreated, &student)
}

func DeleteStudent(ctx *gin.Context) {
	var student entity.Student
	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	database.DB.Delete(&student, id)
	ctx.JSON(http.StatusNoContent, nil)
}

func EditStudent(ctx *gin.Context) {
	var student entity.Student

	id := ctx.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := entity.ValidateStudent(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	ctx.JSON(http.StatusOK, &student)
}

func SearchStudentByCpf(ctx *gin.Context) {
	var student entity.Student
	cpf := ctx.Param("cpf")

	database.DB.Where(&entity.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	ctx.JSON(http.StatusOK, &student)
}
