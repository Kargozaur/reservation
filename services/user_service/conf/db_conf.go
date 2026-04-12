package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConf() *gorm.DB {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("DBUSER"),
		os.Getenv("PASSWORD"),
		os.Getenv("UDBNAME"),
		os.Getenv("PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the db: %v", err)
	}
	return db
}
