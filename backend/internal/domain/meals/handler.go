package meals

import (
	"messify/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GenerateMill(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "Mill generate successfully", nil)
}
func (h *handler) GetAll(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "Mills retrived successfully", nil)
}
func (h *handler) FindById(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "Mill retrived successfully", nil)
}
func (h *handler) Update(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "Mill updated successfully", nil)
}
func (h *handler) Delete(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "Mill deleted successfully", nil)
}