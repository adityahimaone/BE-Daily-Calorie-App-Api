package users

import (
	"Daily-Calorie-App-API/business/users"
	"Daily-Calorie-App-API/controllers"
	_request "Daily-Calorie-App-API/controllers/users/request"
	_response "Daily-Calorie-App-API/controllers/users/response"
	"github.com/labstack/echo/v4"
	"log"
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
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	domainUser, domainPersonalData := _request.ToDomain(&req)
	resp, err := controller.serviceUser.RegisterUser(domainUser, domainPersonalData)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) GetUserById(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := controller.serviceUser.GetUserByID(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) UpdateUser(echoContext echo.Context) error {
	req := _request.User{}
	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	domainUser, domainPersonalData := _request.ToDomain(&req)
	userID := 1
	resp, err := controller.serviceUser.EditUser(userID, domainUser, domainPersonalData)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) DeleteUser(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := controller.serviceUser.DeleteUser(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, resp)
}

func (controller *Controller) GetAllUser(echoContext echo.Context) error {
	resp, err := controller.serviceUser.GetAllUser()
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.FromDomainArray(*resp))
}

func (controller *Controller) LoginUser(echoContext echo.Context) error {
	req := _request.User{}
	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	log.Println(req)
	resp, err := controller.serviceUser.Login(req.Email, req.Password)
	log.Println(resp)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.UserLogin{Token: resp})
}
