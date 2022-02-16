package v1

import (
	"github.com/gocastsian/adamak/adapter/store"
	"github.com/gocastsian/adamak/dto"
	"github.com/gocastsian/adamak/interactor/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func FindUsers(store store.MySQLStore) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req = dto.FindUsersRequest{}

		resp, err := user.New(store).FindUsers(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}
}