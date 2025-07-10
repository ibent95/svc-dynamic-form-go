package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"svc-dynamic-form-go/src/services"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Hello from user controller!",
	})
}
