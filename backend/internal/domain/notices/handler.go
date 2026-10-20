package notices

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

func (h *handler) GenerateNotice(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "notice generated successfully", nil)
}

func (h *handler) GetAll(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "notices retrived successfully", nil)
}

func (h *handler) FindById(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "notice retrived successfully", nil)
}

func (h *handler) Update(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "notice updated successfully", nil)
}

func (h *handler) Delete(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "notice deleted successfully", nil)
}