package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eugenshima/moviori/internal/model"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	handlr AythServiceInterface
}

func NewAuthHandler(handlr AythServiceInterface) *AuthHandler {
	return &AuthHandler{handlr: handlr}
}

type AythServiceInterface interface {
	LoginService(ctx context.Context, auth *model.AuthModel) error
}

func (hnd *AuthHandler) Login(c echo.Context) error {
	var auth *model.AuthModel
	err := c.Bind(&auth)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println("working - ", auth.Login, auth.Password)
	err = hnd.handlr.LoginService(context.Background(), auth)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return nil
}
