package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connectionString := "host=" + dbHost + " user=postgres password=postgres dbname=go-gin port=" + dbPort + " sslmode=disable"
	log.Println("connectionString: ", connectionString)
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("erro ao conectar com o banco de dados.", err)
	}

	err = RunMigrations(DB)

	if err != nil {
		log.Panic("erro ao executar migrações.", err)
	}
}
