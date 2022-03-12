package response

import "Daily-Calorie-App-API/business/meal_plans"

type MealPlans struct {
	UserID             int    `json:"user_id"`
	DietaryPreferences string `json:"dietary_preferences"`
	PlanType           string `json:"plan_type"`
	RangeCalories      string `json:"range_calories"`
	MealPlans          string `json:"meal_plans"`
}

func FromDomain(domain meal_plans.Domain) MealPlans {
	return MealPlans{
		UserID:             domain.UserID,
		DietaryPreferences: domain.DietaryPreferences,
		PlanType:           domain.PlanType,
		RangeCalories:      domain.RangeCalories,
		MealPlans:          domain.MealPlans,
	}
}
