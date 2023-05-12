package main

import (
	"Assigment_6/db"
	"Assigment_6/models"
	"fmt"
	"log"
)

func main() {
	db.ConnectToDb()
	err := db.DB.AutoMigrate(&models.Employee{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Success Migration Table Employee")
}
