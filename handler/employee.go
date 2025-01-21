package handler

import (
	"employee-master/payload"
	"employee-master/response"
	"employee-master/usecase"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeHandler struct {
	employeeUsecase usecase.EmployeeUsecase
}

func NewEmployeeHandler(empUsecase usecase.EmployeeUsecase) EmployeeHandler {
	return EmployeeHandler{
		employeeUsecase: empUsecase,
	}
}

func (e EmployeeHandler) AddEmployee(ctx *gin.Context) {
	var employee payload.EmployeeRequest

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&employee)
	if err != nil {
		response.BadRequestJSON(ctx, "error occured while decoding json :"+err.Error())
		return
	}

	validator := validator.New()
	err = validator.Struct(employee)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid request :"+err.Error())
		return
	}

	employeeData := payload.ToEmployeeModel(employee)

	employeeID, err := e.employeeUsecase.AddEmployee(ctx, employeeData)
	if err != nil {
		response.InternalServerErrorJSON(ctx, "error occured :"+err.Error())
		return
	}

	response.SuccessJSONResponse(ctx, response.IDResponse{ID: employeeID})
}

func (e EmployeeHandler) GetEmployeeByID(ctx *gin.Context) {
	id := ctx.Param("employee_id")
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid employee id")
		return
	}
	employeeDetails, err := e.employeeUsecase.GetEmployee(ctx, employeeID)
	if err != nil {
		if err.Error() == "employee with given id does not exists" {
			response.NotFoundJSON(ctx, err.Error())
			return
		} else {
			response.InternalServerErrorJSON(ctx, err.Error())
			return
		}
	}
	response.SuccessJSONResponse(ctx, response.ToEmployeeResponse(employeeDetails))
}

func (e EmployeeHandler) GetAllEmployees(ctx *gin.Context) {
	employeeData, err := e.employeeUsecase.GetAllEmployee(ctx)
	if err != nil {
		response.InternalServerErrorJSON(ctx, err.Error())
		return
	}
	response.SuccessJSONResponse(ctx, response.ToEmployeeListResponse(employeeData))
}

func (e EmployeeHandler) UpdateEmployee(ctx *gin.Context) {
	id := ctx.Param("employee_id")
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid employee id")
		return
	}

	var empRequest payload.EmployeeRequest

	decoder := json.NewDecoder(ctx.Request.Body)
	err = decoder.Decode(&empRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "error occured while decoding json : "+err.Error())
		return
	}

	validator := validator.New()
	err = validator.Struct(empRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid request data : "+err.Error())
		return
	}

	err = e.employeeUsecase.UpdateEmployee(ctx, employeeID, payload.ToEmployeeModel(empRequest))
	if err != nil {
		response.InternalServerErrorJSON(ctx, "error occured : "+err.Error())
		return
	}

	response.SuccessJSON(ctx, "employee details updated successfully")
}

func (e EmployeeHandler) DeleteEmployee(ctx *gin.Context) {
	id := ctx.Param("employee_id")
	employeeID, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid employee id")
		return
	}
	err = e.employeeUsecase.DeleteEmployee(ctx, employeeID)
	if err != nil {
		if err.Error() == "employee with given id does not exists" {
			response.NotFoundJSON(ctx, err.Error())
			return
		} else {
			response.InternalServerErrorJSON(ctx, err.Error())
			return
		}
	}
	response.SuccessJSON(ctx, "employee deleted successfully")
}
