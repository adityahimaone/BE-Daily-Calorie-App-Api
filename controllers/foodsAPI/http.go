package foodsAPI

import (
	"Daily-Calorie-App-API/business/foodsAPI"
	"Daily-Calorie-App-API/controllers"
	"Daily-Calorie-App-API/controllers/foodsAPI/request"
	"Daily-Calorie-App-API/controllers/foodsAPI/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	foodsAPIService foodsAPI.Service
}

func NewController(foodsAPIService foodsAPI.Service) *Controller {
	return &Controller{
		foodsAPIService: foodsAPIService,
	}
}

func (controller *Controller) GetFoodByName(echoContext echo.Context) error {
	name := echoContext.QueryParam("name")
	resp, err := controller.foodsAPIService.GetFoodByName(name)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, response.FromDomainArray(*resp))
}

func (controller *Controller) GetMealPlan(echoContext echo.Context) error {
	req := request.Food{}
	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	resp, err := controller.foodsAPIService.GetMealPlan(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, response.FromDomainRecipe(*resp))
}
