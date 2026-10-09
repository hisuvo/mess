package routes

import (
	"messify/internal/domain/users"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {

	api := e.Group("/api/v1")

	// welcome api route
	api.GET("", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"status": "ok",
			"message": "Welcome to the Messify API",
			"version": "1.0.0",
			"timestamp": time.Now(),
		})
	})

	// users routes
	users.RegisterRoute(api, db)

	// // Teacher routes
	// teachers.RegisterRoute(api, db)

	// // Student routes
	// students.RegisterRoute(api, db)

	// // Subject routes
	// subjects.RegisterRoute(api, db)
	
}