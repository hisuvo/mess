package users

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	api.POST("/users", handler.Create)
	api.PUT("/users/:id", handler.Update)
	api.DELETE("/users/:id", handler.Delete)
	api.GET("/users/:id", handler.FindByID)
	api.GET("/users/email/:email", handler.FindByEmail)
	api.GET("/users", handler.GetAll)
}