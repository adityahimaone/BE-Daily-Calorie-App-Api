package meal_plans

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/foodsAPI"
)

type serviceMealPlans struct {
	mealplansRepository Repository
	foodsAPIService     foodsAPI.Service
	jwtAuth             *auth.ConfigJWT
}

func NewMealPlansService(mealplansRepository Repository, foodsAPIService foodsAPI.Service, jwtAuth *auth.ConfigJWT) Service {
	return &serviceMealPlans{
		mealplansRepository: mealplansRepository,
		foodsAPIService:     foodsAPIService,
		jwtAuth:             jwtAuth,
	}
}

func (service serviceMealPlans) CreateMealPlans(domain *Domain, domainFoodAPI *foodsAPI.Domain) (*Domain, error) {
	mealplans, _ := service.foodsAPIService.GetMealPlan(domainFoodAPI)
	domain.Lunch = mealplans.Lunch
	domain.Dinner = mealplans.Dinner
	domain.Breakfast = mealplans.Breakfast
	domain.MealTime = mealplans.MealTime
	domain.PlanType = mealplans.PlanType
	domain.DietaryPreferences = mealplans.DietaryPreferences
	result, err := service.mealplansRepository.Insert(domain)
	if err != nil {
		return nil, err
	}
	return result, nil
}
