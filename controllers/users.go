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

type UpdateUserInput struct {
	Name  string `json:"name" validate:"required"`
}

// GET /users
// Find all users
func FindUsers(c echo.Context) error {
	var users []models.User

	if err := models.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

// GET /users/:id
// Find a user
func FindUser(c echo.Context) error {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, user)
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

	if err := models.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// PATCH /users/:id
// Update a user
func UpdateUser(c echo.Context) error {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	// Validate input
	var input = new(UpdateUserInput)
	if err := c.Bind(input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(input); err != nil {
		return err
	}

	user.Name = input.Name

	if err := models.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c echo.Context) error {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	if err := models.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}