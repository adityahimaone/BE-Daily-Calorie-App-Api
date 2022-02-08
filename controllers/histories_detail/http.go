package histories_detail

import (
	"Daily-Calorie-App-API/business/histories_detail"
	"Daily-Calorie-App-API/controllers"
	"Daily-Calorie-App-API/controllers/histories_detail/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Controller struct {
	historiesDetailService histories_detail.Service
}

func NewController(historiesDetailService histories_detail.Service) *Controller {
	return &Controller{
		historiesDetailService: historiesDetailService,
	}
}

func (controller *Controller) DeleteDetailHistory(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	result, err := controller.historiesDetailService.Delete(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, result)
}

func (controller *Controller) GetAllHistoryDetailByHistoriesID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	result, err := controller.historiesDetailService.GetAllHistoriesDetailByHistoriesID(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, response.FromDomainArray(*result))
}
