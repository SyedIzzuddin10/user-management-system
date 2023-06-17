package main

import (
	"userManagementSystem/handler"
	_ "userManagementSystem/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// _ = service.OpenDB()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route
	e.POST("/create-user", handler.CreateUser)
	e.GET("/users", handler.GetUserList)
	e.GET("/user", handler.GetUserById)
	e.PUT("update-user", handler.UpdateUser)
	e.DELETE("delete-user", handler.DeleteUser)

	e.GET("/hello-world", handler.HelloWorld)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
