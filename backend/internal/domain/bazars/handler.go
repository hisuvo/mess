package bazars

import (
	"messify/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	serviec Service
}

func NewHandler(serviec Service) handler {
	return handler{
		serviec: serviec,
	}
}

func (h *handler) BazarCreate(c *echo.Context) error{
	return httpresponse.Success(c, http.StatusOK, "bazar list created successfully", nil)
}

func (h *handler) GetAll(c *echo.Context)  error {
	return httpresponse.Success(c, http.StatusOK, "bazars retrived successfully", nil)
}

func (h *handler) FindById(c *echo.Context)  error {
	return httpresponse.Success(c, http.StatusOK, "bazar retrived successfully", nil)
}

func (h *handler) BazarUpdate(c *echo.Context) error{
	return httpresponse.Success(c, http.StatusOK, "bazar list update successfully", nil)
}

func (h *handler) BazarDelete(c *echo.Context) error{
	return httpresponse.Success(c, http.StatusOK, "bazar list delete successfully", nil)
}