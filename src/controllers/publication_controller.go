package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"svc-dynamic-form-go/src/services"
)

type PublicationController struct {
	Service *services.PublicationService
}

func NewPublicationController(service *services.PublicationService) *PublicationController {
	return &PublicationController{Service: service}
}

func (c *PublicationController) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Hello from publication controller!",
	})
}
