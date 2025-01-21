package usecase

import (
	"context"
	"employee-master/model"
	"errors"
)

var EmployeeIDCounter = 1000

var EmployeeData = map[int]model.Employee{}

type EmployeeUsecases interface {
	AddEmployee(ctx context.Context, Context, employeeDetails model.Employee) error
	GetEmployee(ctx context.Context, emplyeeID int) (model.Employee, error)
	GetAllEmployee(ctx context.Context) ([]model.Employee, error)
	UpdateEmployee(ctx context.Context, employeeDetails model.Employee) error
	DeleteEmployee(ctx context.Context, emplyeeID int) error
}

type EmployeeUsecase struct {
	// Add Dependancies if any
	// eg. database dependancies
}

func NewEmployeeUsecase() EmployeeUsecase {
	return EmployeeUsecase{}
}

func generateEmployeeID() int {
	EmployeeIDCounter = EmployeeIDCounter + 1
	return EmployeeIDCounter
}

func (eu EmployeeUsecase) AddEmployee(ctx context.Context, employee model.Employee) (int, error) {
	id := generateEmployeeID()
	employee.EmployeeID = id
	EmployeeData[id] = employee
	return id, nil
}

func (eu EmployeeUsecase) GetEmployee(ctx context.Context, employeeID int) (model.Employee, error) {
	if empData, exists := EmployeeData[employeeID]; exists {
		return empData, nil
	} else {
		return model.Employee{}, errors.New("employee with given id does not exists")
	}
}

func (eu EmployeeUsecase) GetAllEmployee(ctx context.Context) ([]model.Employee, error) {
	var employees []model.Employee

	for _, val := range EmployeeData {
		employees = append(employees, val)
	}
	return employees, nil
}

func (eu EmployeeUsecase) UpdateEmployee(ctx context.Context, employeeID int, employee model.Employee) error {
	employee.EmployeeID = employeeID
	if _, exists := EmployeeData[employeeID]; exists {
		EmployeeData[employeeID] = employee
		return nil
	} else {
		return errors.New("employee with given id does not exists")
	}
}

func (eu EmployeeUsecase) DeleteEmployee(ctx context.Context, employeeID int) error {
	if _, exists := EmployeeData[employeeID]; exists {
		delete(EmployeeData, employeeID)
		return nil
	} else {
		return errors.New("employee with given id does not exists")
	}
}
