package sql_data

import (
	"context"
	"employee/internal/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type EmployeeSalaryStorer interface {
	GetByEmpId(ctx context.Context, id int) (model.Salary, error)
	Create(ctx context.Context, emp_id int, sal model.Salary) (model.Salary, error)
}

type employeeSalaryStore struct {
	db *gorm.DB
}

func (e employeeSalaryStore) GetByEmpId(ctx context.Context, id int) (model.Salary, error) {
	query := fmt.Sprintf("select * from employee_salary where employee_id=%v", id)
	emp := model.Salary{}
	result := e.db.Raw(query).Scan(&emp)
	if result.Error != nil {
		return model.Salary{}, result.Error
	}
	return emp, nil
}

func (e employeeSalaryStore) Create(ctx context.Context, emp_id int, sal model.Salary) (model.Salary, error) {
	sqlQuery := "INSERT INTO employee_salary (salary, currency, start_date, end_date, employee_id) VALUES (?, ?, ?, ?, ?)"
	result := e.db.Exec(sqlQuery, sal.Salary, sal.Currency, time.Now(), nil, emp_id)
	if result.Error != nil {
		return model.Salary{}, result.Error
	}
	return e.GetByEmpId(ctx, emp_id)
}
func (e employeeSalaryStore) Delete(ctx context.Context, emp_id int) error {
	sqlQuery := "DELETE employee_salary WHERE employee_id=?"

	result := e.db.Exec(sqlQuery, emp_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func NewEmployeeSalary(db *gorm.DB) EmployeeSalaryStorer {
	return &employeeSalaryStore{db: db}
}
