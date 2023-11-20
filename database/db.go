package database

import (
	"log"

	"github.com/joao-antonio-gomes/go-gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connectionString := "host=localhost user=postgres password=postgres dbname=go-gin port=5435 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("erro ao conectar com o banco de dados.")
	}

	DB.AutoMigrate(&models.Student{})
}
