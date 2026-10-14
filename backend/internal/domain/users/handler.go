package users

import (
	"messify/internal/domain/users/dto"
	"messify/internal/httpresponse"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

type Handler interface {
	Create(c *echo.Context) error
	Update(c *echo.Context) error
	Delete(c *echo.Context) error
	FindByID(c *echo.Context) error
	FindByEmail(c *echo.Context) error
	FindAll(c *echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service}
}

func (h *handler) Create(c *echo.Context) error {
	var req *dto.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "Invalid request", nil, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "Validation failed", nil, err.Error())
	}

	result, err := h.service.Create(c.Request().Context(), req)

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "User created failed", nil, err.Error())
	}

	return httpresponse.Success(c, http.StatusCreated, "User created successfully", result)
}

func (h *handler) Update(c *echo.Context) error {
	return nil
}
func (h *handler) Delete(c *echo.Context) error {
	return nil
}
func (h *handler) FindByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return httpresponse.Error(c, http.StatusBadRequest, "Invalid user ID", nil, err.Error())
	}

	result, err := h.service.FindByID(c.Request().Context(), uint(id))

	if err != nil {
		return httpresponse.Error(c, http.StatusInternalServerError, "Failed to find user", nil, err.Error())
	}

	return httpresponse.Success(c, http.StatusOK, "User found successfully", result)
}
func (h *handler) FindByEmail(c *echo.Context) error {
	return nil
}
func (h *handler) FindAll(c *echo.Context) error {
	return c.JSON(http.StatusOK,map[string]any{
		"success":true,
		"message":"Get users successfully",
		"data":[]dto.UserResponse{},
	})
}