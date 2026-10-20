package tenants

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.POST("/tenant", handler.CreateTenant)
	api.GET("/tenants", handler.GetAll)
	api.GET("/tenant/:id", handler.GetById)
	api.PUT("/tenant",handler.Update)
	api.GET("/tenant/:email", handler.FindByEmail)
	api.DELETE("/tenant/:id", handler.Delete)

}