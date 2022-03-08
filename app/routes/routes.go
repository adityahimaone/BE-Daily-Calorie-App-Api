package routes

import (
	auth "Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business"
	controller "Daily-Calorie-App-API/controllers"
	"Daily-Calorie-App-API/controllers/admins"
	"Daily-Calorie-App-API/controllers/foods"
	"Daily-Calorie-App-API/controllers/foodsAPI"
	"Daily-Calorie-App-API/controllers/histories"
	"Daily-Calorie-App-API/controllers/histories_detail"
	"Daily-Calorie-App-API/controllers/meal_plans"
	"Daily-Calorie-App-API/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type HandlerList struct {
	UserController            users.Controller
	JWTMiddleware             middleware.JWTConfig
	FoodController            foods.Controller
	AdminController           admins.Controller
	HistoriesController       histories.Controller
	HistoriesDetailController histories_detail.Controller
	FoodAPIController         foodsAPI.Controller
	MealPlansController       meal_plans.Controller
}

func (controller *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")

	//Admin Endpoint
	group.POST("/admin/login", controller.AdminController.LoginAdmin)
	group.POST("/admin/register", controller.AdminController.RegisterAdmin)

	//User Endpoint
	group.POST("/users/login", controller.UserController.LoginUser)
	group.POST("/users/register", controller.UserController.RegisterUser)
	group.GET("/users/:id", controller.UserController.GetUserById, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.PUT("/users/edit/:id", controller.UserController.UpdateUser, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.DELETE("/users/:id", controller.UserController.DeleteUser, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationAdmin())
	group.GET("/users/", controller.UserController.GetAllUser, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationAdmin())

	//Food Endpoint
	group.GET("/foods/:id", controller.FoodController.GetFoodByID, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.GET("/foods", controller.FoodController.GetAllFood, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.GET("/foods/", controller.FoodController.GetFoodByName, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.POST("/foods/add", controller.FoodController.CreateFood, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.PUT("/foods/edit/:id", controller.FoodController.UpdateFood, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.DELETE("/foods/:id", controller.FoodController.DeleteFood, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationAdmin())

	//History Endpoint
	group.POST("/histories/add", controller.HistoriesController.CreateHistories, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.POST("/histories/water", controller.HistoriesController.CreateWater, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.POST("/histories/automatic", controller.HistoriesController.CreateHistoriesFromAPI, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.GET("/histories/last", controller.HistoriesController.GetLastHistoriesByUserID, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.GET("/histories/list", controller.HistoriesController.GetAllHistoriesByUserID, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.GET("/histories/:id", controller.HistoriesController.GetHistoriesByID, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.DELETE("/histories/:id", controller.HistoriesController.DeleteHistoriesDetail, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())

	//Histories Detail Endpoint
	group.GET("/histories_detail/all/:id", controller.HistoriesDetailController.GetAllHistoryDetailByHistoriesID, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.DELETE("/histories_detail/:id", controller.HistoriesDetailController.DeleteDetailHistory, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())

	//Foods API Endpoint
	group.GET("/foods/api/", controller.FoodAPIController.GetFoodByName, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())
	group.POST("/open-api/meal-plan", controller.FoodAPIController.GetMealPlan, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUserAndAdmin())

	//Meal Plans Endpoint
	group.POST("/meal-plan", controller.MealPlansController.CreateMealPlans, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())
	group.GET("/meal-plan", controller.MealPlansController.GetLastMealPlans, middleware.JWTWithConfig(controller.JWTMiddleware), RoleValidationUser())

}

func RoleValidationUser() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(echoContext echo.Context) error {
			claims := auth.GetUser(echoContext)

			if claims.Role == "user" {
				return hf(echoContext)
			} else {
				return controller.NewErrorResponse(echoContext, http.StatusForbidden, business.ErrUnAuthorized)
			}
		}
	}
}

func RoleValidationAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(echoContext echo.Context) error {
			claims := auth.GetUser(echoContext)

			if claims.Role == "admin" {
				return hf(echoContext)
			} else {
				return controller.NewErrorResponse(echoContext, http.StatusForbidden, business.ErrUnAuthorized)
			}
		}
	}
}

func RoleValidationUserAndAdmin() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(echoContext echo.Context) error {
			claims := auth.GetUser(echoContext)

			if claims.Role == "admin" || claims.Role == "user" {
				return hf(echoContext)
			} else {
				return controller.NewErrorResponse(echoContext, http.StatusForbidden, business.ErrUnAuthorized)
			}
		}
	}
}
