package users

import (
	"Daily-Calorie-App-API/businesses/users"
	"Daily-Calorie-App-API/controllers/users/request"
	_response "Daily-Calorie-App-API/controllers/users/response"
	"Daily-Calorie-App-API/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	serviceUser users.Service
}

func NewController(serviceUser users.Service) *Controller {
	return &Controller{
		serviceUser: serviceUser,
	}
}

func (controller *Controller) RegisterUser(echoContext echo.Context) error {
	req := request.User{}
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Register", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}

	resp1, resp2, err := controller.serviceUser.RegisterUser(req.ToDomain())
	if err != nil {
		response := helpers.APIResponse("Failed Register", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Register", http.StatusOK, "Success", _response.FromDomain(*resp1, *resp2))
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *Controller) GetUserById(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp1, resp2, err := controller.serviceUser.FindByID(id)
	if err != nil {
		response := helpers.APIResponse("Failed Find User", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Find User", http.StatusOK, "Success", _response.FromDomain(*resp1, *resp2))
	return echoContext.JSON(http.StatusOK, response)
}
