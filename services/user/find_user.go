package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gocastsian/adamak/protocol"
	"github.com/labstack/echo/v4"
)

var FindUserService findUserService

type findUserService struct{}

func (s *findUserService) Process(ctx context.Context, req FindUserRequest) (res FindUserResponse, err error) {
	if err = doesUserExist(ctx, req.ID); err != nil {
		return
	}

	user, err := storage.GetUser(ctx, req.ID)
	if err != nil {
		return
	}
	res = FindUserResponse{User: user}
	return
}

func (s *findUserService) ServeHTTPEcho(c echo.Context) error {
	return findUserEcho(c)
}

type FindUserRequest struct {
	ID uint `json:"-"`
}

type FindUserResponse struct {
	User protocol.User `json:"user"`
}

func findUserEcho(c echo.Context) error {
	idStr := c.Param("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var req = FindUserRequest{ID: uint(userID)}
	res, err := FindUserService.Process(c.Request().Context(), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
	return nil
}
