package response

import "employee-master/model"

type EmployeeResponse struct {
	Name        string  `json:"name"`
	EmployeeID  int     `json:"employee_id"`
	Designation string  `json:"designation"`
	Department  string  `json:"department"`
	Salary      float32 `json:"salary"`
}

func ToEmployeeResponse(emp model.Employee) EmployeeResponse {
	return EmployeeResponse{
		EmployeeID:  emp.EmployeeID,
		Name:        emp.Name,
		Designation: emp.Designation,
		Department:  emp.Department,
		Salary:      emp.Salary,
	}
}

func ToEmployeeListResponse(employees []model.Employee) []EmployeeResponse {
	var employeesResponse []EmployeeResponse
	for _, emp := range employees {
		employeesResponse = append(employeesResponse, EmployeeResponse{
			EmployeeID:  emp.EmployeeID,
			Name:        emp.Name,
			Designation: emp.Designation,
			Department:  emp.Department,
			Salary:      emp.Salary,
		})
	}
	return employeesResponse
}
