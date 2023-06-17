package handler

import (
	"context"
	"net/http"
	"userManagementSystem/dto"
	"userManagementSystem/service"

	"github.com/labstack/echo/v4"
)

// var userService = service.NewUserDatabase()

func CreateUser(c echo.Context) error {
	// name := c.QueryParam("name")
	// email := c.QueryParam("email")
	// password := c.QueryParam("password")
	// userService := service.NewUserDatabase()
	// userService.CreateUser(name, email, password)

	userService := service.NewUserService(c, context.Background())
	user, err := userService.CreateUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, dto.Display{
		Message: "User created successfully!",
		Data:    user,
	})
}
