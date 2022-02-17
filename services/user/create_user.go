package user

import (
	"context"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gocastsian/adamak/protocol"
	"github.com/labstack/echo/v4"
)

var CreateUserService createUserService

type createUserService struct{}

func (s *createUserService) Process(ctx context.Context, req CreateUserRequest) (res CreateUserResponse, err error) {
	err = s.Validate(req)
	if err != nil {
		return
	}

	var user = protocol.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	res.User, err = storage.CreateUser(ctx, user)
	return
}

func (s *createUserService) Validate(req CreateUserRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required),
	)
}

func (s *createUserService) ServeHTTPEcho(c echo.Context) error {
	return createUserEcho(c)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	User protocol.User `json:"user"`
}

func createUserEcho(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := CreateUserService.Process(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
	return nil
}
