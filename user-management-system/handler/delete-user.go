package handler

import (
	"context"
	"net/http"
	"userManagementSystem/dto"
	"userManagementSystem/service"

	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	// id := c.QueryParam("id")

	// userService := service.NewUserDatabase()
	// err := userService.DeleteUser(id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }

	userService := service.NewUserService(c, context.Background())
	user, err := userService.DeleteUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, dto.Display{
		Message: "User has been deleted!",
		Data:    user,
	})
}
