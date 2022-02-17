package user

import (
	"context"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gocastsian/adamak/protocol"
	"github.com/labstack/echo/v4"
)

var UpdateUserService updateUserService

type updateUserService struct{}

func (s *updateUserService) Process(ctx context.Context, req UpdateUserRequest) (res UpdateUserResponse, err error) {
	err = s.Validate(ctx, req)
	if err != nil {
		return
	}
	err = doesUserExist(ctx, req.ID)
	if err != nil {
		return
	}

	var user protocol.User
	user, err = storage.GetUser(ctx, req.ID)
	if err != nil {
		return
	}
	user.Name = req.Name

	res.User, err = storage.UpdateUser(ctx, user)
	return
}

func (s *updateUserService) ServeHTTPEcho(c echo.Context) error {
	return updateUserEcho(c)
}

func (s *updateUserService) Validate(ctx context.Context, req UpdateUserRequest) error {
	return validation.ValidateStruct(&req, validation.Field(&req.Name, validation.Required))
}

type UpdateUserRequest struct {
	ID   uint   `json:"-"`
	Name string `json:"name"`
}

type UpdateUserResponse struct {
	User protocol.User `json:"user"`
}

func updateUserEcho(c echo.Context) error {
	var idStr = c.Param("id")
	var userID, err = strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	req.ID = uint(userID)

	res, err := UpdateUserService.Process(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
	return nil
}
