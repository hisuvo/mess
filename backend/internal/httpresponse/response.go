package httpresponse

import (
	"github.com/labstack/echo/v5"
)


type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    any `json:"data,omitempty"`
	Errors   any `json:"errors,omitempty"`
}

func Success(c *echo.Context, status int, message string, data any)error{
	return c.JSON(status, Response{
		Success: true,
		Message: message,
		Data: data,
	})
}

func Error(c *echo.Context,status int, message string, data any, errors any)error{
	return c.JSON(status, Response{
		Success: false,
		Message: message,
		Data: data,
		Errors: errors,
	})
}