package user

import (
	"context"
	"net/http"

	"github.com/gocastsian/adamak/protocol"
	"github.com/labstack/echo/v4"
)

var FindUsersService findUsersService

type findUsersService struct{}

func (s *findUsersService) Process(ctx context.Context) (res FindUsersResponse, err error) {
	res.Users, err = storage.FindUsers(ctx)
	return
}

func (s *findUsersService) ServeHTTPEcho(c echo.Context) error {
	return findUsersEcho(c)
}

type FindUsersResponse struct {
	Users []protocol.User `json:"users"`
}

func findUsersEcho(c echo.Context) error {
	res, err := FindUsersService.Process(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
	return nil
}
