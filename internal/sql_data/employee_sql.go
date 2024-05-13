package sql_data

import (
	"context"
	"employee/internal/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type EmployeeStorer interface {
	Get(ctx context.Context, id int) (model.Employee, error)
	Create(ctx context.Context, employee model.Employee) (model.Employee, error)
	Update(ctx context.Context, id int, employee model.Employee) (model.Employee, error)
	Delete(ctx context.Context, id int) error
	GetEmployeesPage(page int, pageSize int) ([]model.Employee, error)
}

type employeeStore struct {
	db *gorm.DB
}

func (e employeeStore) GetEmployeesPage(page int, pageSize int) ([]model.Employee, error) {
	var employees []model.Employee
	offset := (page - 1) * pageSize

	// Query employees with pagination
	result := e.db.Offset(offset).Limit(pageSize).Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}

	return employees, nil
}

func (e employeeStore) Get(ctx context.Context, id int) (model.Employee, error) {
	query := fmt.Sprintf("select * from employee where id=%v", id)
	emp := model.Employee{}
	result := e.db.Raw(query).Scan(&emp)
	if result.Error != nil {
		return model.Employee{}, result.Error
	}
	return emp, nil
}

func (e employeeStore) Create(ctx context.Context, employee model.Employee) (model.Employee, error) {
	sqlQuery := "INSERT INTO employee (name, designation, created_at, updated_at) VALUES (?, ?, ?, ?)"

	// Execute the raw SQL query with parameters
	result := e.db.Exec(sqlQuery, employee.Name, employee.Designation, time.Now(), time.Now())
	if result.Error != nil {
		return model.Employee{}, result.Error
	}
	var lastInsertedID uint
	e.db.Raw("SELECT LAST_INSERT_ID()").Scan(&lastInsertedID)

	return e.Get(ctx, int(lastInsertedID))
}

func (e employeeStore) Update(ctx context.Context, id int, employee model.Employee) (model.Employee, error) {
	result := e.db.Model(&model.Employee{}).Where("id = ?", id).Update("designation", employee.Designation)
	if result.Error != nil {
		return model.Employee{}, result.Error
	}
	return e.Get(ctx, id)
}

func (e employeeStore) Delete(ctx context.Context, id int) error {
	sqlQuery := "DELETE FROM employee_insurance WHERE employee_id=?"

	result := e.db.Exec(sqlQuery, id)
	if result.Error != nil {
		return result.Error
	}

	sqlQuery = "DELETE FROM employee_salary WHERE employee_id=?"
	result = e.db.Exec(sqlQuery, id)
	if result.Error != nil {
		return result.Error
	}
	result = e.db.Delete(&model.Employee{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewEmployeeStore(db *gorm.DB) EmployeeStorer {
	return &employeeStore{
		db: db,
	}
}
