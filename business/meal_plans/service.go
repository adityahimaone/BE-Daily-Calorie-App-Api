package meal_plans

import (
	"Daily-Calorie-App-API/app/middleware/auth"
)

type serviceMealPlans struct {
	mealplansRepository Repository
	jwtAuth             *auth.ConfigJWT
}

func NewMealPlansService(mealplansRepository Repository, jwtAuth *auth.ConfigJWT) Service {
	return &serviceMealPlans{
		mealplansRepository: mealplansRepository,
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
