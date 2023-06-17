package handler

import (
	"fmt"
	"net/http"
	"userManagementSystem/dto"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	// id := c.QueryParam("id")

	// userService := service.NewUserDatabase()
	// err := userService.DeleteUser(id)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, err)
	// }
	fmt.Println("Hello World!")

	return c.JSON(http.StatusOK, dto.Display{
		Message: "This is testing endpoint for pi-server",
		Data:    "Hello World!",
	})
}
