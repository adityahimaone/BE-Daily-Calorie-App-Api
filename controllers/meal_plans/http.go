package meal_plans

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/meal_plans"
	"Daily-Calorie-App-API/controllers"
	"Daily-Calorie-App-API/controllers/meal_plans/request"

	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	mealplansService meal_plans.Service
}

func NewController(mealplansService meal_plans.Service) *Controller {
	return &Controller{
		mealplansService: mealplansService,
	}
}

func (controller *Controller) CreateMealPlans(echoContext echo.Context) error {
	req := request.MealPlans{}
	if err := echoContext.Bind(&req); err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	user := auth.GetUser(echoContext)
	req.UserID = user.ID
	domainMealPlans, domainFoodsAPI := request.ToDomain(&req)
	resp, err := controller.mealplansService.CreateMealPlans(domainMealPlans, domainFoodsAPI)
	if err != nil {
		return controllers.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}
	return echoContext.JSON(http.StatusCreated, resp)
}
