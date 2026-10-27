package tenants

import (
	"messify/internal/domain/tenants/dto"
	"messify/internal/httpresponse"
	"messify/internal/pkg/validator"
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
	var tenant dto.RequestCreateTenant

	if err := c.Bind(&tenant); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "invalied tenant field request", nil, err.Error())
	}

	if err := c.Validate(&tenant); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "tenant validation request failed", nil, err.Error())
	}

	if err := h.service.CreateTenant(c.Request().Context(), &tenant); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "tenant created request failed", nil, err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "Tenant create successfully", nil)
}

func (h *handler) GetAll(c *echo.Context) error {
	tenants, err := h.service.GetAll(c.Request().Context())

	if err != nil {
		return httpresponse.Error(c, http.StatusOK, "Tenant retrived successfully", nil, err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "Tenant retrived successfully", tenants)
}

func (h *handler) GetById(c *echo.Context) error {

	id , err := validator.ParseUintID(c.Param("id"))

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalid tenant ID", nil, err.Error())
	}

	tenant, err := h.service.GetById(c.Request().Context(), id)


	if err != nil {
		return httpresponse.Error(c, http.StatusOK,"your tenant get by id successfullyD", nil, err.Error())
	}

	return httpresponse.Success(c, http.StatusOK,"your tenant get by id successfullyD", tenant)
}

func (h *handler) Delete(c *echo.Context) error {
	id, err := validator.ParseUintID(c.Param("id"))

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalid tenant ID", nil, err.Error())
	}

	result, err := h.service.SoftDelete(c.Request().Context(), id)

	if err != nil {
		return httpresponse.Error(c, http.StatusNotFound, "tenant delete failed", nil, err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "tenant deleted successfullyD", result)
}

func (h *handler) Update(c *echo.Context) error {
	id, err := validator.ParseUintID(c.Param("id"))

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest,"Invalid tenant ID", nil, err.Error())
	}

	var req dto.RequestUpdateTenant

	// Id parse in reqest inside
	req.ID = id

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "invalied tenant field request", nil, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "tenant validation request failed", nil, err.Error())
	}

	response, err := h.service.Update(c.Request().Context(), id, &req);

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "tenant created request failed", nil, err.Error())
	}
	
	return httpresponse.Success(c, http.StatusOK, "tenant update successfully", response)
}