package routes

import (
	"Daily-Calorie-App-API/controllers/foods"
	"Daily-Calorie-App-API/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	UserController users.Controller
	JWTMiddleware  middleware.JWTConfig
	FoodController foods.Controller
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")

	//User Endpoint
	group.POST("/users/register", handler.UserController.RegisterUser)
	group.GET("/users/:id", handler.UserController.GetUserById)
	group.PUT("/users/edit_user", handler.UserController.UpdateUser, middleware.JWTWithConfig(handler.JWTMiddleware))
	//group.DELETE("/users/:id", handler.UserController.DeleteUser)
	group.POST("/users/login", handler.UserController.LoginUser)
	group.GET("/users/", handler.UserController.GetAllUsers)

	//Food Endpoint
	group.GET("/foods/:id", handler.FoodController.GetFoodByID)
	group.GET("/foods/", handler.FoodController.GetAllFood)
	group.POST("/foods/add_food", handler.FoodController.CreateFood)
	group.PUT("/foods/edit_food/:id", handler.FoodController.UpdateFood)
	group.DELETE("/foods/:id", handler.FoodController.DeleteFood)
}
