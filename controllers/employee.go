package controllers

import (
	"Assigment_6/db"
	"Assigment_6/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	Code          int         `json:"code"`
	Status        string      `json:"status"`
	Message       string      `json:"message"`
	MessageDetail string      `json:"message_detail"`
	Data          interface{} `json:"data"`
}

var body struct {
	FullName string
	Email    string
	Age      int
	Division string
}
var employee models.Employee

func PostEmp(c *gin.Context) {
	c.Bind(&body)
	post := models.Employee{
		FullName: body.FullName,
		Email:    body.Email,
		Age:      body.Age,
		Division: body.Division,
	}
	result := db.DB.Create(&post)
	if result.Error != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Fail",
			Message:       "Fail To Add Employee",
			MessageDetail: result.Error.Error(),
			Data:          nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := ApiResponse{
		Code:          http.StatusCreated,
		Status:        "Success",
		Message:       "Success To Add Employee",
		MessageDetail: "",
		Data:          post,
	}
	c.JSON(http.StatusCreated, response)
}

func GetAllEmp(c *gin.Context) {
	var post []models.Employee
	result := db.DB.Find(&post)
	if result.Error != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Fail",
			Message:       "Fail To Get All Employee",
			MessageDetail: result.Error.Error(),
			Data:          nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := ApiResponse{
		Code:          http.StatusOK,
		Status:        "Success",
		Message:       "Success To Get All Employee",
		MessageDetail: "",
		Data:          post,
	}
	c.JSON(http.StatusAccepted, response)

}

func GetEmpById(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	employee := models.Employee{ID: uint(num)}
	result := db.DB.Take(&employee)
	if result.Error != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Fail",
			Message:       "Fail To Get Employee",
			MessageDetail: result.Error.Error(),
			Data:          nil,
		}

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := ApiResponse{
		Code:          http.StatusOK,
		Status:        "Success",
		Message:       "Success To Get Employee",
		MessageDetail: "",
		Data:          employee,
	}
	c.JSON(http.StatusOK, response)

}

func UpdateEmp(c *gin.Context) {
	c.Bind(&body)
	id := c.Param("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	db.DB.First(&employee, num)
	err := db.DB.Model(&employee).Updates(&body).Error
	if err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Fail",
			Message:       "Fail To Update Employee",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := ApiResponse{
		Code:          http.StatusOK,
		Status:        "Success",
		Message:       "Success To Update Employee",
		MessageDetail: "",
		Data:          employee,
	}
	c.JSON(http.StatusOK, response)
}

func DeleteEmp(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.ParseUint(id, 10, 64)
	err := db.DB.Clauses(clause.Returning{}).Delete(&employee, num).Error
	if err != nil {
		response := ApiResponse{
			Code:          http.StatusBadRequest,
			Status:        "Fail",
			Message:       "Fail To Update Employee",
			MessageDetail: err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := ApiResponse{
		Code:          http.StatusOK,
		Status:        "Success",
		Message:       "Success To Update Employee",
		MessageDetail: "",
		Data:          employee,
	}
	c.JSON(http.StatusOK, response)

}
