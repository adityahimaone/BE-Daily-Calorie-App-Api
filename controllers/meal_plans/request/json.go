package request

import (
	"Daily-Calorie-App-API/business/foodsAPI"
	"Daily-Calorie-App-API/business/meal_plans"
)

type MealPlans struct {
	UserID             int    `json:"user_id"`
	MealTime           string `json:"meal_time"`
	DietaryPreferences string `json:"dietary_preferences"`
	PlanType           string `json:"plan_type"`
	RangeCalories      string `json:"range_calories"`
}

func ToDomain(request *MealPlans) (*meal_plans.Domain, *foodsAPI.Domain) {
	return &meal_plans.Domain{
			UserID:             request.UserID,
			MealTime:           request.MealTime,
			DietaryPreferences: request.DietaryPreferences,
			PlanType:           request.PlanType,
			RangeCalories:      request.RangeCalories,
		}, &foodsAPI.Domain{
			MealTime:           request.MealTime,
			DietaryPreferences: request.DietaryPreferences,
			PlanType:           request.PlanType,
			RangeCalories:      request.RangeCalories,
		}
}
