package company

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start() {
	initdb()
	migrate()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/employees", getEmployees)
	e.GET("/employees/:id", getEmployee)
	e.POST("/employees", createEmployee)
	e.PUT("/employees/:id", updateEmployee)
	e.PATCH("/employees/:id", patchEmployee)
	e.DELETE("/employees/:id", deleteEmployee)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
