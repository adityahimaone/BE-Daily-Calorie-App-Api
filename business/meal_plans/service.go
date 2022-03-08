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

func (service serviceMealPlans) CreateMealPlans(domain *Domain) (*Domain, error) {
	result, err := service.mealplansRepository.Insert(domain)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service serviceMealPlans) GetLastMealPlans(userID int) (*Domain, error) {
	result, err := service.mealplansRepository.GetLastMealPlans(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
