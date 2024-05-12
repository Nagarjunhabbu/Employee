package service

import (
	"context"
	"employee/internal/model"
	"errors"

	"gorm.io/gorm"
)

type MockSalaryStorer struct {
}

func (e MockSalaryStorer) GetByEmpId(ctx context.Context, id int) (model.Salary, error) {
	if id == 0 {
		return model.Salary{}, errors.New("id is empty")
	}

	r := model.Salary{
		ID:         2,
		Salary:     120000,
		Currency:   "IDR",
		EmployeeId: 1,
	}
	return r, nil
}

func (e MockSalaryStorer) Create(ctx context.Context, emp_id int, sal model.Salary) (model.Salary, error) {
	return model.Salary{}, nil
}

func NewEmployeeSalary(db *gorm.DB) MockSalaryStorer {
	return MockSalaryStorer{}
}
