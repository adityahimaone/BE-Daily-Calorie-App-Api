package foods

import (
	"Daily-Calorie-App-API/business/foods"
	"Daily-Calorie-App-API/controllers"
	"Daily-Calorie-App-API/controllers/foods/request"
	_response "Daily-Calorie-App-API/controllers/foods/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	foodsService foods.Service
}

func NewController(serviceFood foods.Service) *Controller {
	return &Controller{
		foodsService: serviceFood,
	}
}

func (controller *Controller) GetAllFood(echoContext echo.Context) error {
	resp, err := controller.foodsService.GetAllFood()
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.FromDomainArray(*resp))
}

func (controller *Controller) GetFoodByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := controller.foodsService.GetFoodByID(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) GetFoodByName(echoContext echo.Context) error {
	name := echoContext.QueryParam("name")
	resp, err := controller.foodsService.GetFoodByName(name)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) CreateFood(echoContext echo.Context) error {
	req := request.Food{}

	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	resp, err := controller.foodsService.AddFood(req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessCreatedResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) UpdateFood(echoContext echo.Context) error {
	req := request.Food{}

	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	foodID, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	getData, err := controller.foodsService.GetFoodByID(foodID)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	domainFood := req.ToDomain()
	resp, err := controller.foodsService.EditFood(foodID, domainFood)
	resp.ID = getData.ID
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) DeleteFood(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))

	resp, err := controller.foodsService.DeleteFood(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, resp)
}
