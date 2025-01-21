package payload

import "employee-master/model"

type EmployeeRequest struct {
	Name        string  `json:"name" validate:"required"`
	Designation string  `json:"designation" validate:"required"`
	Department  string  `json:"department" validate:"required"`
	Salary      float32 `json:"salary" validate:"required"`
}

func ToEmployeeModel(request EmployeeRequest) model.Employee {
	return model.Employee{
		Name:        request.Name,
		Designation: request.Designation,
		Department:  request.Department,
		Salary:      request.Salary,
	}
}
