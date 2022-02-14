package controllers

import (
	"github.com/gocastsian/adamak/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateUserInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// POST /users
// Create new user
func CreateUser(c echo.Context) error {
	// Validate input
	var input = new(CreateUserInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return err
	}

	// Create user
	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}
	models.DB.Create(&user)

	return c.JSON(http.StatusOK, user)
}
