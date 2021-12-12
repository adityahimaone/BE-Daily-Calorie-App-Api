package routes

import (
	"Daily-Calorie-App-API/controllers/users"
	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	UserController users.Controller
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")

	//User Endpoint
	group.POST("/users/register", handler.UserController.RegisterUser)
	group.GET("/users/:id", handler.UserController.GetUserById)
}
