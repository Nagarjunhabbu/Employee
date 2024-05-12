package service

import (
	"context"
	"employee/internal/model"
	"errors"
	"testing"
)

func TestGetUserAPI(t *testing.T) {
	service := employeeService{
		Data:     MockSql{},
		SalStore: MockSalaryStorer{},
		InsStore: MockInsurance{},
	}

	tc1 := struct {
		id       int
		expected *model.Employee
		err      error
	}{
		id:       0,
		expected: nil,
		err:      errors.New("id is empty"),
	}

	_, err := service.GetEmployee(context.Background(), tc1.id)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc1.err)
	}

	tc2 := struct {
		id       int
		expected *model.Employee
		err      error
	}{
		id: 1,
		expected: &model.Employee{
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
		},
		err: nil,
	}
	data, err := service.GetEmployee(context.Background(), tc2.id)
	if err != nil {
		t.Errorf("failed test case, expected no error got %v  ", err)
	}
	if data.ID != tc2.expected.ID {
		t.Errorf("failed test case, expected %v  got %v  ", tc2.expected.ID, data.ID)

	}
}

func TestCreateUserAPI(t *testing.T) {
	service := employeeService{
		Data:     MockSql{},
		SalStore: MockSalaryStorer{},
		InsStore: MockInsurance{},
	}
	emp := model.Employee{
		ID:          1,
		Name:        "",
		Designation: "Manager",
		Insurance: &model.Insurance{
			ID:          1,
			InsuranceID: "SKR124",
			InsuranceNo: "INS3490098898864",
			EmployeeId:  1,
		},
		Salary: &model.Salary{
			ID:         2,
			Salary:     10000,
			Currency:   "IDR",
			EmployeeId: 1,
		},
	}

	tc1 := struct {
		emp      model.Employee
		expected *model.Employee
		err      error
	}{
		emp:      emp,
		expected: nil,
		err:      errors.New("employee name is empty"),
	}

	_, err := service.CreateEmployee(context.Background(), tc1.emp)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc1.err)
	}

	emp.Name = "Nags"
	tc2 := struct {
		emp      model.Employee
		expected *model.Employee
		err      error
	}{
		emp: emp,
		expected: &model.Employee{
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
		},
		err: nil,
	}
	data, err := service.CreateEmployee(context.Background(), tc2.emp)
	if err != nil {
		t.Errorf("failed test case, expected no error got %v  ", err)
	}
	if data.ID != tc2.expected.ID {
		t.Errorf("failed test case, expected %v  got %v  ", tc2.expected.ID, data.ID)

	}
}

func TestUpdateUserAPI(t *testing.T) {
	service := employeeService{
		Data:     MockSql{},
		SalStore: MockSalaryStorer{},
		InsStore: MockInsurance{},
	}

	emp := model.Employee{
		ID:          1,
		Designation: "CTO",
		Insurance:   &model.Insurance{},
	}

	tc1 := struct {
		id       int
		emp      model.Employee
		expected *model.Employee
		err      error
	}{
		id:       0,
		emp:      emp,
		expected: nil,
		err:      errors.New("invalid employee id"),
	}

	_, err := service.UpdateEmployee(context.Background(), tc1.id, tc1.emp)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc1.err)
	}

	tc2 := struct {
		id       int
		emp      model.Employee
		expected *model.Employee
		err      error
	}{
		id:  1,
		emp: emp,
		expected: &model.Employee{
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
		},
		err: nil,
	}
	data, err := service.UpdateEmployee(context.Background(), tc2.id, tc2.emp)
	if err != nil {
		t.Errorf("failed test case, expected no error got %v  ", err)
	}
	if data.Designation != tc2.expected.Designation {
		t.Errorf("failed test case, expected %v  got %v  ", tc2.expected.ID, data.ID)

	}
}

func TestDeleteUserAPI(t *testing.T) {
	service := employeeService{
		Data:     MockSql{},
		SalStore: MockSalaryStorer{},
		InsStore: MockInsurance{},
	}

	tc1 := struct {
		id       int
		expected *model.Employee
		err      error
	}{
		id:       0,
		expected: nil,
		err:      errors.New("id is empty"),
	}

	err := service.DeleteEmployee(context.Background(), tc1.id)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc1.err)
	}

	tc2 := struct {
		id  int
		err error
	}{
		id:  1,
		err: nil,
	}
	err = service.DeleteEmployee(context.Background(), tc2.id)
	if err != nil {
		t.Errorf("failed test case, expected no error got %v  ", err)
	}
}

func TestListEmployee(t *testing.T) {
	service := employeeService{
		Data:     MockSql{},
		SalStore: MockSalaryStorer{},
		InsStore: MockInsurance{},
	}

	tc1 := struct {
		page     int
		pageSize int
		expected []*model.Employee
		err      error
	}{
		page:     0,
		pageSize: 1,
		expected: nil,
		err:      errors.New("id is empty"),
	}

	_, err := service.ListEmployee(context.Background(), tc1.page, tc1.pageSize)
	if err == nil {
		t.Errorf("failed test case, expected %v got nil", tc1.err)
	}

	tc2 := struct {
		page     int
		pageSize int
		expected []model.Employee
		err      error
	}{
		page:     1,
		pageSize: 2,
		expected: []model.Employee{{
			ID:          1,
			Name:        "nags",
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
		}},
		err: nil,
	}
	data, err := service.ListEmployee(context.Background(), tc2.page, tc2.pageSize)
	if err != nil {
		t.Errorf("failed test case, expected no error got %v  ", err)
	}
	if data[0].Name != tc2.expected[0].Name {
		t.Errorf("failed test case, expected %v  got %v  ", tc2.expected[0].Name, data[0].Name)

	}
}
