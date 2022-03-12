package foodsAPI

import (
	"Daily-Calorie-App-API/app/middleware/auth"
)

type serviceFoodAPI struct {
	foodAPIRepository Repository
	jwtAuth           *auth.ConfigJWT
}

func NewFoodAPIService(repositoryFoodAPI Repository, jwtAuth *auth.ConfigJWT) Service {
	return &serviceFoodAPI{
		foodAPIRepository: repositoryFoodAPI,
		jwtAuth:           jwtAuth,
	}
}

func (service serviceFoodAPI) GetFoodByName(name string) (*[]Domain, error) {
	result, err := service.foodAPIRepository.GetFoodByName(name)
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}

func (service serviceFoodAPI) GetMealPlan(food *Domain) (*Domain, error) {
	food.MealTime = "Breakfast"
	breakfast, err := service.foodAPIRepository.GetMealPlan(food)
	if err != nil {
		return &Domain{}, err
	}
	food.Breakfast = *breakfast
	food.MealTime = "Lunch"
	lunch, err := service.foodAPIRepository.GetMealPlan(food)
	if err != nil {
		return &Domain{}, err
	}
	food.Lunch = *lunch
	food.MealTime = "Dinner"
	dinner, err := service.foodAPIRepository.GetMealPlan(food)
	if err != nil {
		return &Domain{}, err
	}
	food.Dinner = *dinner
	return food, nil
}
