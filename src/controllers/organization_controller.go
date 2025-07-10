package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"svc-dynamic-form-go/src/services"
)

type OrganizationController struct {
	Service *services.OrganizationService
	UserService *services.UserService
}

func NewOrganizationController(
	service *services.OrganizationService,
	userService *services.UserService,
) *OrganizationController {
	return &OrganizationController{
		Service: service,
		UserService: userService,
	}
}

func (c *OrganizationController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Hello from organization controller!",
	})
}

func (c *OrganizationController) Store(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Hello from store organization controller!",
	})
}
