package database

import (
	"github.com/joao-antonio-gomes/go-gin-rest-api/application/domain/student/entity"
	"gorm.io/gorm"
)

func RunMigrations(gdb *gorm.DB) error {
	err := gdb.AutoMigrate(&entity.Student{})

	return err
}
