package servers

import (
	"fmt"
	"messify/internal/config"
	"messify/internal/routes"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
  validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func Srart(db *gorm.DB, config *config.Config) error {
	e := echo.New()
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	// welcome api route
	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"status": "ok",
			"message": "Welcome to the Messify API",
			"timestamp": time.Now(),
		})
	})

	// middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())


	// api routes
	routes.RegisterRoutes(e, db)

	// cors 
	err := e.Start(":"+ config.Port)

	if err != nil {
		fmt.Println("server one connected failed: ", err.Error())
		return err
	}

	fmt.Println("Server started successfully")
	return nil

}