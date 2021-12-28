package foods

import (
	"Daily-Calorie-App-API/businesses/foods"
	"Daily-Calorie-App-API/controllers/foods/request"
	"Daily-Calorie-App-API/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	serviceFood foods.Service
}

func NewController(serviceFood foods.Service) *Controller {
	return &Controller{
		serviceFood: serviceFood,
	}
}

func (controller *Controller) GetAllFood(echoContext echo.Context) error {
	resp, err := controller.serviceFood.GetFoods()
	if err != nil {
		response := helpers.APIResponse("Failed Get All Foods", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Get All Food", http.StatusOK, "Success", *resp)
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *Controller) GetFoodByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := controller.serviceFood.GetFood(id)
	if err != nil {
		response := helpers.APIResponse("Failed Get Food By ID", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Get Food By ID", http.StatusOK, "Success", *resp)
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *Controller) CreateFood(echoContext echo.Context) error {
	req := request.Food{}
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Create Food", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	domainFood, domainNutrition := request.ToDomain(req)
	resp, err := controller.serviceFood.AddFood(domainFood, domainNutrition)
	if err != nil {
		response := helpers.APIResponse("Failed Create Food", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Create Food", http.StatusOK, "Success", *resp)
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *Controller) UpdateFood(echoContext echo.Context) error {
	req := request.Food{}
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Update Food", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	domainFood, _ := request.ToDomain(req)
	resp, err := controller.serviceFood.EditFood(domainFood.ID, domainFood)
	if err != nil {
		response := helpers.APIResponse("Failed Update Food", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Update Food", http.StatusOK, "Success", *resp)
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *Controller) DeleteFood(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := controller.serviceFood.DeleteFood(id)
	if err != nil {
		response := helpers.APIResponse("Failed Delete Food", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Delete Food", http.StatusOK, "Success", *resp)
	return echoContext.JSON(http.StatusOK, response)
}
