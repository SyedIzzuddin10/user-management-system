package handler

import (
	"context"
	"net/http"
	"userManagementSystem/dto"
	"userManagementSystem/service"

	"github.com/labstack/echo/v4"
)

func GetUserList(c echo.Context) error {

	// userService := service.NewUserDatabase()
	// userList := userService.GetUserList()

	userService := service.NewUserService(c, context.Background())
	userList, err := userService.GetUserList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if userList == nil {
		return c.JSON(http.StatusBadRequest, "Database is empty!")
	}

	return c.JSON(http.StatusOK, dto.Display{
		Message: "User list in database",
		Data:    userList,
	})
}
