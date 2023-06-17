package handler

import (
	"context"
	"net/http"
	"userManagementSystem/dto"
	"userManagementSystem/service"

	"github.com/labstack/echo/v4"
)

func GetUserById(c echo.Context) error {
	// id := c.QueryParam("id")

	// userService := service.NewUserDatabase()
	// user, err := userService.GetUserById(id)

	userService := service.NewUserService(c, context.Background())
	user, err := userService.GetUserById()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, dto.Display{
		Message: "User data found",
		Data:    user,
	})
}
