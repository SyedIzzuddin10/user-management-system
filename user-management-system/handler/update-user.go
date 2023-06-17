package handler

import (
	"context"
	"net/http"
	"userManagementSystem/dto"
	"userManagementSystem/service"

	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
	// id := c.QueryParam("id")
	// name := c.QueryParam("name")
	// email := c.QueryParam("email")
	// password := c.QueryParam("password")

	// userService := service.NewUserDatabase()
	// user, err := userService.UpdateUser(id, name, email, password)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	userService := service.NewUserService(c, context.Background())
	user, err := userService.UpdateUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, dto.Display{
		Message: "User updated successfully",
		Data:    user,
	})
}
