package tenants

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

func (h *handler) CreateTenant(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "Tenant create successfully", nil)
}

func (h *handler) GetAll(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK, "Tenant retrived successfully", nil)
}

func (h *handler) GetById(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK,"your tenant get by id successfullyD", nil)
}

func (h *handler) FindByEmail(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK,"Tenants get by email", nil)
}

func (h *handler) Delete(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK,"tenant deleted successfullyD", nil)
}

func (h *handler) Update(c *echo.Context) error {
	return httpresponse.Success(c, http.StatusOK,"tenant update successfully", nil)
}