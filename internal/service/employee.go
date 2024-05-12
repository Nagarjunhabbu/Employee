package service

import (
	"context"
	"employee/internal/model"
	"employee/internal/sql_data"
	"errors"
	"fmt"
	"log"
	"sync"
)

type EmployeeService interface {
	GetEmployee(ctx context.Context, id int) (model.Employee, error)
	CreateEmployee(ctx context.Context, employee model.Employee) (model.Employee, error)
	UpdateEmployee(ctx context.Context, id int, employee model.Employee) (model.Employee, error)
	DeleteEmployee(ctx context.Context, id int) error
	ListEmployee(ctx context.Context, page int, pageSize int) ([]model.Employee, error)
}

type employeeService struct {
	Data     sql_data.EmployeeStorer
	SalStore sql_data.EmployeeSalaryStorer
	InsStore sql_data.EmployeeInsuranceStorer
}

func (e employeeService) ListEmployee(ctx context.Context, page int, pageSize int) ([]model.Employee, error) {
	data, err := e.Data.GetEmployeesPage(page, pageSize)
	if err != nil {
		return nil, err
	}
	var returnResp []model.Employee
	for _, v := range data {
		resp, _ := e.GetEmployee(ctx, v.ID)
		returnResp = append(returnResp, resp)
	}
	return returnResp, nil
}

func (e employeeService) GetEmployee(ctx context.Context, id int) (model.Employee, error) {
	m, err := e.Data.Get(ctx, id)
	if err != nil {
		log.Println("error in getting employee", err)
		return model.Employee{}, err
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(m *model.Employee) {

		val, err := e.SalStore.GetByEmpId(ctx, m.ID)
		if err != nil {
			fmt.Println("error in getting salary", err)
		}
		m.Salary = &val
		wg.Done()
	}(&m)

	go func(m *model.Employee) {
		ins, err := e.InsStore.GetByEmpId(ctx, m.ID)
		if err != nil {
			fmt.Println("error in getting insurance", err)
		}
		m.Insurance = &ins
		wg.Done()
	}(&m)
	wg.Wait()

	return m, nil
}

func (e employeeService) CreateEmployee(ctx context.Context, req model.Employee) (model.Employee, error) {
	if req.Name == "" {
		log.Println("employee name is empty")
		return model.Employee{}, errors.New("employee name is empty")

	}

	if req.Designation == "" {
		log.Println("employee designation is empty")
		return model.Employee{}, errors.New("employee designation is empty")
	}

	emp, err := e.Data.Create(ctx, req)
	if err != nil {
		log.Println("error in creating employee", err)
		return model.Employee{}, err
	}
	wg := &sync.WaitGroup{}
	if req.Salary != nil {
		wg.Add(1)
		go func(emp *model.Employee, req model.Employee) {
			sal, _ := e.SalStore.Create(ctx, emp.ID, *req.Salary)
			emp.Salary = &sal
			wg.Done()
		}(&emp, req)
	}
	if req.Insurance != nil {
		wg.Add(1)
		go func(emp *model.Employee, req model.Employee) {
			ins, _ := e.InsStore.Create(ctx, emp.ID, *req.Insurance)
			emp.Insurance = &ins
			wg.Done()
		}(&emp, req)
	}
	wg.Wait()

	return emp, nil
}

func (e employeeService) UpdateEmployee(ctx context.Context, id int, employee model.Employee) (model.Employee, error) {
	employee, err := e.Data.Update(ctx, id, employee)
	if err != nil {
		log.Println("error in updating an employee details", err)
		return model.Employee{}, err
	}
	return employee, nil
}

func (e employeeService) DeleteEmployee(ctx context.Context, id int) error {
	return e.Data.Delete(ctx, id)
}

func NewEmployeeService(storer sql_data.EmployeeStorer, salStore sql_data.EmployeeSalaryStorer, ins sql_data.EmployeeInsuranceStorer) EmployeeService {
	return &employeeService{Data: storer, SalStore: salStore, InsStore: ins}
}
