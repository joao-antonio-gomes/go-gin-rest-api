package repository

import (
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/domain/student/entity"
	"github.com/joao-antonio-gomes/go-gin-rest-api/database"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (r *StudentRepository) FindAll() ([]entity.Student, error) {
	var students []entity.Student
	err := r.DB.Find(&students).Error
	return students, err
}

func (r *StudentRepository) FindById(id string) (entity.Student, error) {
	var student entity.Student
	err := r.DB.First(&student, id).Error
	return student, err
}

func (r *StudentRepository) Create(student entity.Student) (entity.Student, error) {
	database.DB.Create(&student)
	return student, nil
}

func (r *StudentRepository) Update(student entity.Student) (entity.Student, error) {
	database.DB.Save(&student)
	return student, nil
}

func (r *StudentRepository) Delete(id string) error {
	var student entity.Student
	err := r.DB.Delete(&student, id).Error
	return err
}

func (r *StudentRepository) SearchStudentByCpf(cpf string) (entity.Student, error) {
	var student entity.Student
	err := r.DB.Where("cpf = ?", cpf).First(&student).Error
	return student, err
}
