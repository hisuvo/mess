package deposits

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoute(api *echo.Group, db *gorm.DB){
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	api.GET("/deposits",handler.GetAll)
	api.GET("/deposit",handler.FindById)
	api.POST("/deposit",handler.GenerateDeposit)
	api.PUT("/deposit/:id",handler.Update)
	api.DELETE("/deposit/:id",handler.Delete)
}