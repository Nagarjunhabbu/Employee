package main

import (
	"employee/internal/controller"
	"employee/internal/service"
	"employee/internal/sql_data"
	"fmt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db := getDb()
	// all initialisation
	storer := sql_data.NewEmployeeStore(db)
	salStorer := sql_data.NewEmployeeSalary(db)
	insStorer := sql_data.NewEmployeeInsurance(db)
	service := service.NewEmployeeService(storer, salStorer, insStorer)
	ctrl := controller.NewEmployeeController(service)

	e.GET("/v1/employee/:id", ctrl.GetEmployees)
	e.GET("/v1/employee", ctrl.ListEmployees)
	e.POST("/v1/employee", ctrl.CreateEmployee)
	e.PATCH("/v1/employee/:id", ctrl.UpdateEmployee)
	e.DELETE("/v1/employee/:id", ctrl.DeleteEmployee)
	e.Logger.Fatal(e.Start(":8000"))
}

func getDb() *gorm.DB {
	var db *gorm.DB
	var err error

	// Wait for MySQL to become available
	for {
		db, err = gorm.Open(mysql.Open("root:root123@tcp(mysql:3306)/employeedb"), &gorm.Config{})
		if err == nil {
			break // MySQL is available, break the loop
		}

		fmt.Println("MySQL is not available yet. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	// MySQL is now available, start your application logic here
	fmt.Println("MySQL is now available. Starting the application...")
	return db
}
