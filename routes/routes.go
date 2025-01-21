package routes

import (
	"employee-master/handler"
	"employee-master/usecase"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	employeeUsecase := usecase.NewEmployeeUsecase()

	employeeHandler := handler.NewEmployeeHandler(employeeUsecase)

	empRoutes := r.Group("employees")
	empRoutes.GET("/", employeeHandler.GetAllEmployees)
	empRoutes.GET("/:employee_id", employeeHandler.GetEmployeeByID)
	empRoutes.POST("/", employeeHandler.AddEmployee)
	empRoutes.PUT("/:employee_id", employeeHandler.UpdateEmployee)
	empRoutes.DELETE("/:employee_id", employeeHandler.DeleteEmployee)
}
