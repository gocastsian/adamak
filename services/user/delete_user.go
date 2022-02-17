package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var DeleteUserService deleteUserService

type deleteUserService struct{}

func (s *deleteUserService) Process(ctx context.Context, req DeleteUserRequest) (err error) {
	err = doesUserExist(ctx, req.ID)
	if err != nil {
		return
	}

	err = storage.DeleteUser(ctx, req.ID)
	return
}

func (s *deleteUserService) ServeHTTPEcho(c echo.Context) error {
	return deleteUserEcho(c)
}

type DeleteUserRequest struct {
	ID uint `json:"-"`
}

func deleteUserEcho(c echo.Context) error {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var req = DeleteUserRequest{ID: uint(userID)}
	err = DeleteUserService.Process(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
