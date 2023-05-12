package db

import (
	"Assigment_6/initializers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDb() error {
	initializers.LoadEnv()
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To Connect")
		return err
	}
	return nil
}
