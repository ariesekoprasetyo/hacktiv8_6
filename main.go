package main

import (
	"Assigment_6/controllers"
	"Assigment_6/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.ConnectToDb()
	if err == nil {
		fmt.Println("Success Connect To Database")
	}
	r := gin.Default()

	r.POST("/employee", controllers.PostEmp)
	r.GET("/employee", controllers.GetAllEmp)
	r.GET("/employee/:id", controllers.GetEmpById)
	r.PUT("/employee/:id", controllers.UpdateEmp)
	r.DELETE("/employee/:id", controllers.DeleteEmp)

	r.Run()
}
