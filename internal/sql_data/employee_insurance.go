package sql_data

import (
	"context"
	"employee/internal/model"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type EmployeeInsuranceStorer interface {
	GetByEmpId(ctx context.Context, id int) (model.Insurance, error)
	Create(ctx context.Context, emp_id int, employee model.Insurance) (model.Insurance, error)
}

type employeeInsuranceStore struct {
	db *gorm.DB
}

func (e employeeInsuranceStore) GetByEmpId(ctx context.Context, id int) (model.Insurance, error) {
	fmt.Println("inside insurance store")
	query := fmt.Sprintf("select * from employee_insurance where employee_id=%v", id)
	emp := model.Insurance{}
	result := e.db.Raw(query).Scan(&emp)
	if result.Error != nil {
		return model.Insurance{}, result.Error
	}
	fmt.Println(emp)
	return emp, nil
}

func (e employeeInsuranceStore) Create(ctx context.Context, emp_id int, ins model.Insurance) (model.Insurance, error) {
	sqlQuery := "INSERT INTO employee_insurance (insurance_id, insurance_no, insurance_exp, employee_id) VALUES (?, ?, ?, ?)"

	result := e.db.Exec(sqlQuery, ins.InsuranceID, ins.InsuranceNo, time.Now(), emp_id)
	if result.Error != nil {
		return model.Insurance{}, result.Error
	}
	return e.GetByEmpId(ctx, emp_id)
}

func NewEmployeeInsurance(db *gorm.DB) EmployeeInsuranceStorer {
	return &employeeInsuranceStore{db: db}
}
