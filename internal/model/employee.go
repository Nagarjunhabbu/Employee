package model

type Employee struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name"`
	Designation string     `json:"designation"`
	Salary      *Salary    `json:"salary"`
	Insurance   *Insurance `json:"insurance"`
}

func (e Employee) TableName() string {
	return "employee"
}

type Salary struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	Salary     float64 `json:"salary"`
	Currency   string  `json:"currency"`
	EmployeeId int
}

func (s Salary) TableName() string {
	return "employee_salary"
}

type Insurance struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	InsuranceID string `json:"insurance_id"`
	InsuranceNo string `json:"insurance_no"`
	EmployeeId  int    `json:"-"`
}

func (i Insurance) TableName() string {
	return "employee_insurance"
}
