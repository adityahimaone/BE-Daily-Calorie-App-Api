package request

import "Daily-Calorie-App-API/business/foodsAPI"

type Food struct {
	MealTime           string `json:"mealTime"`
	DietaryPreferences string `json:"dietaryPreferences"`
	PlanType           string `json:"planType"`
	RangeCalories      string `json:"rangeCalories"`
}

func (request *Food) ToDomain() *foodsAPI.Domain {
	return &foodsAPI.Domain{
		MealTime:           request.MealTime,
		DietaryPreferences: request.DietaryPreferences,
		PlanType:           request.PlanType,
		RangeCalories:      request.RangeCalories,
	}
}
