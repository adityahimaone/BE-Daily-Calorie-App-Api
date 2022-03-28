package histories

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/histories"
	"Daily-Calorie-App-API/controllers"
	"Daily-Calorie-App-API/controllers/histories/request"
	"Daily-Calorie-App-API/controllers/histories/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
	user := auth.GetUser(echoContext)
	req.UserID = user.ID
	resp, err := controller.historiesService.CreateHistories(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessCreatedResponse(echoContext, response.FromDomain(*resp))
}

func (controller *Controller) CreateWater(echoContext echo.Context) error {
	req := request.HistoriesWater{}
	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	user := auth.GetUser(echoContext)
	req.UserID = user.ID
	resp, err := controller.historiesService.CreateWater(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, response.FromDomain(*resp))
}

func (controller *Controller) GetLastHistoriesByUserID(echoContext echo.Context) error {
	user := auth.GetUser(echoContext)
	resp, err := controller.historiesService.GetLastHistoryByUserID(user.ID)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, response.FromDomain(*resp))
	//return controllers.NewSuccessResponse(echoContext, resp)
}

func (controller *Controller) GetAllHistoriesByUserID(echoContext echo.Context) error {
	user := auth.GetUser(echoContext)
	resp, err := controller.historiesService.GetAllHistoryByUserID(user.ID)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, response.FromDomainArray(*resp))
}

func (controller *Controller) GetHistoriesByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := controller.historiesService.GetHistoriesByID(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, response.FromDomain(*resp))
}

func (controller *Controller) CreateHistoriesFromAPI(echoContext echo.Context) error {
	req := request.HistoriesFood{}

	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	user := auth.GetUser(echoContext)
	req.UserID = user.ID
	domainFoods, domainHistories := request.ToDomain(&req)
	resp, err := controller.historiesService.CreateHistoriesFromAPI(domainHistories, domainFoods)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessCreatedResponse(echoContext, response.FromDomain(*resp))
}

func (controller *Controller) DeleteHistoriesDetail(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := controller.historiesService.DeleteHistoriesDetail(id)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(echoContext, resp)
}
