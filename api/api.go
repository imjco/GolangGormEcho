package api

import (
	"github.com/labstack/echo/v4"
	"profile/handlers"
)

func InitializeRoutes(e *echo.Echo) {
	g := e.Group("/api/v1")
	g.GET("/get-customer-list", handlers.GetCustomerList)
	g.GET("/get-customer-list/:customerid", handlers.GetCustomerbyID)
	g.POST("/add-customer", handlers.AddCustomer)
	g.POST("/update-customer", handlers.AddCustomer)

}
