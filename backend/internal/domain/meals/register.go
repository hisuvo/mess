package meals

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.POST("/meal", handler.GenerateMill)
	api.GET("/meals", handler.GetAll)
	api.GET("/meal/:id", handler.FindById)
	api.PUT("/meal/:id", handler.Update)
	api.DELETE("/meal/:id", handler.Delete)
}