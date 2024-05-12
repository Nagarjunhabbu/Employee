package service

import (
	"context"
	"employee/internal/model"
	"errors"
)

type MockSql struct {
}

func (e MockSql) GetEmployeesPage(page int, pageSize int) ([]model.Employee, error) {
	if page == 0 {
		return []model.Employee{}, errors.New("invalid page number")
	}

	r := []model.Employee{{
		ID:          1,
		Name:        "Nags",
		Designation: "CTO",
		Insurance: &model.Insurance{
			ID:          1,
			InsuranceID: "SKR123",
			InsuranceNo: "INS349009889889",
			EmployeeId:  1,
		},
		Salary: &model.Salary{
			ID:         2,
			Salary:     120000,
			Currency:   "IDR",
			EmployeeId: 1,
		},
	}, {
		ID:          2,
		Name:        "Ganesh",
		Designation: "Manager",
		Insurance: &model.Insurance{
			ID:          1,
			InsuranceID: "SKR124",
			InsuranceNo: "INS349009889459",
			EmployeeId:  2,
		},
		Salary: &model.Salary{
			ID:         3,
			Salary:     3000000,
			Currency:   "IDR",
			EmployeeId: 2,
		},
	}}
	return r, nil
}

func (e MockSql) Get(ctx context.Context, id int) (model.Employee, error) {
	if id == 0 {
		return model.Employee{}, errors.New("id is empty")
	}

	r := model.Employee{
		ID:          1,
		Name:        "nags",
		Designation: "CEO",
		Insurance: &model.Insurance{
			ID:         1,
			EmployeeId: 1,
		},
	}
	return r, nil
}

func (e MockSql) Create(ctx context.Context, employee model.Employee) (model.Employee, error) {
	if employee.Name == "" {
		return model.Employee{}, errors.New("employee Name is empty")
	}

	if employee.Designation == "" {
		return model.Employee{}, errors.New("employee Designation is empty")
	}

	r := model.Employee{
		ID:          1,
		Name:        "Nags",
		Designation: "CEO",
		Insurance: &model.Insurance{
			ID:          1,
			InsuranceID: "SKR123",
			InsuranceNo: "INS349009889889",
			EmployeeId:  1,
		},
		Salary: &model.Salary{
			ID:         2,
			Salary:     120000,
			Currency:   "IDR",
			EmployeeId: 1,
		},
	}
	return r, nil
}

func (e MockSql) Update(ctx context.Context, id int, employee model.Employee) (model.Employee, error) {

	if id == 0 {
		return model.Employee{}, errors.New("id is empty")
	}

	r := model.Employee{
		ID:          1,
		Name:        "Nags",
		Designation: "CTO",
		Insurance: &model.Insurance{
			ID:          1,
			InsuranceID: "SKR123",
			InsuranceNo: "INS349009889889",
			EmployeeId:  1,
		},
		Salary: &model.Salary{
			ID:         2,
			Salary:     120000,
			Currency:   "IDR",
			EmployeeId: 1,
		},
	}
	return r, nil
}

func (e MockSql) Delete(ctx context.Context, id int) error {
	if id == 0 {
		return errors.New("id is empty")
	}
	return nil
}
