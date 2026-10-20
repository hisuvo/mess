package bazars

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group , db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.GET("/bazars", handler.GetAll)
	api.GET("/bazar/:id", handler.FindById)
	api.POST("/bazar", handler.BazarCreate)
	api.PUT("/bazar/:id", handler.BazarUpdate)
	api.DELETE("/bazar/:id", handler.BazarDelete)
}