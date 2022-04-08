package meal_plans

import (
	"Daily-Calorie-App-API/business/meal_plans"
	"gorm.io/gorm"
)

type repositoryMealPlans struct {
	DB *gorm.DB
}

func NewRepositoryMealPlans(db *gorm.DB) meal_plans.Repository {
	return &repositoryMealPlans{
		DB: db,
	}
}

func (repository repositoryMealPlans) Insert(domain *meal_plans.Domain) (*meal_plans.Domain, error) {
	recordMealPlans := fromDomain(*domain)
	if err := repository.DB.Save(&recordMealPlans).Error; err != nil {
		return nil, err
	}
	result := recordMealPlans.toDomain()
	return &result, nil
}

func (repository repositoryMealPlans) GetLastMealPlans(userID int) (*meal_plans.Domain, error) {
	recordMealPlans := MealPlans{}
	if err := repository.DB.Where("user_id = ?", userID).Last(&recordMealPlans).Error; err != nil {
		return nil, err
	}
	result := recordMealPlans.toDomain()
	return &result, nil
}
