package company

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func initdb() {
	dsn := "host=localhost user=postgres password=postgresql dbname=companydb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established")
}
func migrate() {
	db.AutoMigrate(&Employee{})
}

type Employee struct {
	Id         uint   `gorm:"primaryKey" json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Salary     int    `json:"salary"`
	Experience int    `json:"experience"`
	CreatedAt  string `json:"created_at"`
}

func getEmployees(c echo.Context) error {
	var employees []Employee
	db.Find(&employees)
	return c.JSON(http.StatusOK, employees)
}

func getEmployee(c echo.Context) error {
	id := c.Param("id")
	var employee Employee

	if err := db.First(&employee, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Employee not found"})
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, employee)
}

func createEmployee(c echo.Context) error {
	reqBody := new(Employee)
	if err := c.Bind(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.Create(&reqBody).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, reqBody)
}

func updateEmployee(c echo.Context) error {
	id := c.Param("id")
	var employee Employee

	if err := db.First(&employee, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Employee not found"})
	}

	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.Save(&employee).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, employee)
}

func patchEmployee(c echo.Context) error {
	id := c.Param("id")
	var employee Employee

	if err := db.First(&employee, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Employee not found"})
	}

	type UpdateEmployee struct {
		Name       *string `json:"name"`
		Age        *int    `json:"age"`
		Salary     *int    `json:"salary"`
		Experience *int    `json:"experience"`
		CreatedAt  *string `json:"created_at"`
	}

	var updateEmployee UpdateEmployee
	if err := c.Bind(&updateEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if updateEmployee.Name != nil {
		employee.Name = *updateEmployee.Name
	}
	if updateEmployee.Age != nil {
		employee.Age = *updateEmployee.Age
	}
	if updateEmployee.Salary != nil {
		employee.Salary = *updateEmployee.Salary
	}
	if updateEmployee.Experience != nil {
		employee.Experience = *updateEmployee.Experience
	}
	if updateEmployee.CreatedAt != nil {
		employee.CreatedAt = *updateEmployee.CreatedAt
	}

	// Save the updated employee
	if err := db.Save(&employee).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, employee)
}

func deleteEmployee(c echo.Context) error {
	id := c.Param("id")
	if err := db.Delete(&Employee{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Employee has been deleted",
		"id":      id,
	})
}
