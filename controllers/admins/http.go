package admins

import (
	"Daily-Calorie-App-API/business/admins"
	"Daily-Calorie-App-API/controllers"
	_request "Daily-Calorie-App-API/controllers/admins/request"
	_response "Daily-Calorie-App-API/controllers/admins/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	serviceAdmin admins.Service
}

func NewController(serviceAdmin admins.Service) *Controller {
	return &Controller{
		serviceAdmin: serviceAdmin,
	}
}

func (controller *Controller) RegisterAdmin(echoContext echo.Context) error {
	req := _request.Admin{}
	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	domainAdmin := _request.ToDomain(&req)
	resp, err := controller.serviceAdmin.Register(domainAdmin)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessCreatedResponse(echoContext, _response.FromDomain(*resp))
}

func (controller *Controller) LoginAdmin(echoContext echo.Context) error {
	req := _request.Admin{}
	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	resp, err := controller.serviceAdmin.Login(req.Username, req.Password)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(echoContext, _response.AdminLogin{
		Token: resp,
	})
}
