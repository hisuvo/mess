package notices

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.GET("/notices", handler.GetAll)
	api.GET("/notice/:id", handler.FindById)
	api.POST("/notice", handler.GenerateNotice)
	api.PUT("/notice/:id", handler.Update)
	api.DELETE("/notice/:id", handler.Delete)
}