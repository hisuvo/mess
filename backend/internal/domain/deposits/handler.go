package deposits

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

func (h *handler) GenerateDeposit(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "deposit Created successfully", nil)
}

func (h *handler) GetAll(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "deposits retrived successfully", nil)
}

func (h *handler) FindById(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "deposit retrived successfully", nil)
}

func (h *handler) Update(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "deposit updated successfully", nil)
}

func (h *handler) Delete(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "deposit deleted successfully", nil)
}
