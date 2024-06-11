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
	LoginService(ctx context.Context, login *model.UserModel) error
	SignupService(ctx context.Context, auth *model.UserModel) error
}

func (hnd *AuthHandler) Login(c echo.Context) error {
	var user *model.UserModel
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println("working - ", user.Login, user.Password)
	err = hnd.handlr.LoginService(context.Background(), user)
	if err != nil {
		return c.String(http.StatusNotAcceptable, "Wrong password")
	}
	return c.String(http.StatusOK, "Password is right")
}

func (hnd *AuthHandler) Signup(c echo.Context) error {
	var user *model.UserModel
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println("working - ", user.Login, user.Password)
	err = hnd.handlr.SignupService(context.Background(), user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return nil
}
