package v1

import (
	"github.com/gocastsian/adamak/adapter/store"
	"github.com/gocastsian/adamak/contract"
	"github.com/gocastsian/adamak/dto"
	"github.com/gocastsian/adamak/interactor/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateUser(store store.MySQLStore, validator contract.ValidateCreateUser) echo.HandlerFunc {
	return func (c echo.Context) error {

		var req = dto.CreateUserRequest{}
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := validator(req); err != nil {
			return err
		}

		resp, err := user.New(store).CreateUser(c.Request().Context(), req)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, resp)
	}}
