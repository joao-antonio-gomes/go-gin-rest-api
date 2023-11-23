package seeds

import (
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/domain/student/entity"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func createStudent(db *gorm.DB, firstProp string, secondProp string) error {
	return db.Create(&entity.Student{
		Name: firstProp,
		CPF:  secondProp,
	}).Error
}

func All() []Seed {
	return []Seed{
		{
			Name: "Example-student-1",
			Run: func(db *gorm.DB) error {
				return createStudent(db, "Joãozinho", "12345678901")
			},
		},
	}
}
