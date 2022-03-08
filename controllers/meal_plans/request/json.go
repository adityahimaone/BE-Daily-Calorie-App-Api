package request

import (
	"Daily-Calorie-App-API/business/meal_plans"
)

type MealPlans struct {
	UserID             int    `json:"user_id"`
	DietaryPreferences string `json:"dietary_preferences"`
	PlanType           string `json:"plan_type"`
	RangeCalories      string `json:"range_calories"`
	MealPlans          string `json:"meal_plans"`
}

func ToDomain(request *MealPlans) *meal_plans.Domain {
	return &meal_plans.Domain{
		UserID:             request.UserID,
		DietaryPreferences: request.DietaryPreferences,
		PlanType:           request.PlanType,
		RangeCalories:      request.RangeCalories,
		MealPlans:          request.MealPlans,
	}
}
