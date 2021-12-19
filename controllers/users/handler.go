package users

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/businesses/users"
	_request "Daily-Calorie-App-API/controllers/users/request"
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
	req := _request.User{}
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Register", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domainUser, domainPersonalData := _request.ToDomain(&req)
	resp1, resp2, err := controller.serviceUser.RegisterUser(domainUser, domainPersonalData)
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

func (controller *Controller) UpdateUser(echoContext echo.Context) error {
	req := _request.User{}
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Update User", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domainUser, domainPersonalData := _request.ToDomain(&req)
	userAuth := auth.GetUser(echoContext)
	userID := userAuth.ID
	resp1, resp2, err := controller.serviceUser.EditUser(userID, domainUser, domainPersonalData)
	if err != nil {
		response := helpers.APIResponse("Failed Update User", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Update User", http.StatusOK, "Success", _response.FromDomain(*resp1, *resp2))
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *Controller) LoginUser(echoContext echo.Context) error {
	req := _request.User{}
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Login User", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := controller.serviceUser.Login(req.Email, req.Password)
	if err != nil {
		response := helpers.APIResponse("Failed Login User", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Login User", http.StatusOK, "Success", _response.UserLogin{Token: resp})
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *Controller) GetAllUsers(echoContext echo.Context) error {
	resp, err := controller.serviceUser.GetAllUsers()
	if err != nil {
		response := helpers.APIResponse("Failed Get All Users", http.StatusInternalServerError, "Error", err)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Get All Users", http.StatusOK, "Success", _response.FromDomainArray(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
