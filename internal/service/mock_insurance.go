package service

import (
	"context"
	"employee/internal/model"
	"errors"
)

type MockInsurance struct {
}

func (e MockInsurance) GetByEmpId(ctx context.Context, id int) (model.Insurance, error) {
	if id == 0 {
		return model.Insurance{}, errors.New("id is empty")
	}

	r := model.Insurance{
		ID:          1,
		InsuranceID: "SKR123",
		InsuranceNo: "INS349009889889",
		EmployeeId:  1,
	}
	return r, nil
}

func (e MockInsurance) Create(ctx context.Context, emp_id int, ins model.Insurance) (model.Insurance, error) {
	return model.Insurance{}, nil
}

func NewEmployeeInsurance() MockInsurance {
	return MockInsurance{}
}
