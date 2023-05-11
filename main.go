package main

import (
	"Assigment_6/db"
	"Assigment_6/initializers"
)

func init() {
	initializers.LoadEnv()
	db.ConnectToDb()
}

func main() {
}
