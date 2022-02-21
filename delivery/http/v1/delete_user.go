package v1

import (
	"github.com/gocastsian/adamak/adapter/store"
	"github.com/gocastsian/adamak/contract"
	"github.com/gocastsian/adamak/dto"
	"github.com/gocastsian/adamak/interactor/user"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func DeleteUser(store store.MySQLStore, validator contract.ValidateDeleteUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		userID, err := strconv.Atoi(idStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var req = dto.DeleteUserRequest{ID: uint(userID)}

		if err := validator(c.Request().Context(), req); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		if _, err := user.New(store).DeleteUser(c.Request().Context(), req); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}