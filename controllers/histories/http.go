package histories

import (
	"Daily-Calorie-App-API/business/histories"
	"Daily-Calorie-App-API/controllers"
	"Daily-Calorie-App-API/controllers/histories/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	historiesService histories.Service
}

func NewController(serviceHistories histories.Service) *Controller {
	return &Controller{historiesService: serviceHistories}
}

func (controller *Controller) CreateHistories(echoContext echo.Context) error {
	req := request.Histories{}

	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	resp, err := controller.historiesService.CreateHistories(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, resp)
}
