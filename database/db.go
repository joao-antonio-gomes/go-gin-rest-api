package database

import (
	"log"
	"os"

	"github.com/joao-antonio-gomes/go-gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	connectionString := "host=" + dbHost + " user=postgres password=postgres dbname=go-gin port=5432 sslmode=disable"
	log.Println("connectionString: ", connectionString)
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("erro ao conectar com o banco de dados.", err)
	}

	DB.AutoMigrate(&models.Student{})
}
