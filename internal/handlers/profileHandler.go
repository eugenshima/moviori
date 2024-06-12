package handlers

import (
	"context"
	"net/http"

	"github.com/eugenshima/moviori/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	handlr AuthServiceInterface
}

func NewAuthHandler(handlr AuthServiceInterface) *AuthHandler {
	return &AuthHandler{handlr: handlr}
}

type AuthServiceInterface interface {
	LoginService(ctx context.Context, login *model.UserModel) (*model.FullUserModel, error)
	SignupService(ctx context.Context, auth *model.UserModel) error
}

func (hnd *AuthHandler) Login(c echo.Context) error {
	var user *model.UserModel
	err := c.Bind(&user)
	if err != nil {
		logrus.Errorf("Login: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to bind info")
	}

	userinfo, err := hnd.handlr.LoginService(context.Background(), user)
	if err != nil {
		logrus.Errorf("Login: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong login/password")
	}
	return c.JSON(http.StatusOK, userinfo)
}

func (hnd *AuthHandler) Signup(c echo.Context) error {
	var user *model.UserModel
	err := c.Bind(&user)
	if err != nil {
		logrus.Errorf("Signup: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to bind info")
	}

	err = hnd.handlr.SignupService(context.Background(), user)
	if err != nil {
		logrus.Errorf("Signup: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed to sign up")
	}
	return c.String(http.StatusOK, "Account has been created")
}
