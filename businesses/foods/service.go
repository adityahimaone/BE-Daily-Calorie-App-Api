package foods

import (
	"Daily-Calorie-App-API/businesses/nutritioninfo"
	"log"
)

type serviceFoods struct {
	foodRepository      Repository
	nutritionRepository nutritioninfo.Repository
}

func NewService(repositoryFood Repository, repositoryNutrition nutritioninfo.Repository) Service {
	return &serviceFoods{
		foodRepository:      repositoryFood,
		nutritionRepository: repositoryNutrition,
	}
}

func (service serviceFoods) AddFood(food *Domain, nutrition *nutritioninfo.Domain) (*Domain, error) {
	nutritionInfo, err := service.nutritionRepository.Create(nutrition)
	if err != nil {
		return nil, err
	}
	food.NutritionInfoID = nutritionInfo.ID
	food.Calories = nutritionInfo.Calories
	food.Protein = nutritionInfo.Protein
	food.Carbs = nutritionInfo.Carbs
	food.Fat = nutritionInfo.Fat
	food.ServingSize = nutritionInfo.ServingSize
	log.Println(nutritionInfo.ID, "nutrition info id")
	log.Println(food.Calories, food.Carbs, food.Protein, food.Fat)
	result, err := service.foodRepository.InsertFood(food)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service serviceFoods) EditFood(id int, food *Domain) (*Domain, error) {
	result, err := service.foodRepository.UpdateFood(id, food)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service serviceFoods) DeleteFood(id int) (*Domain, error) {
	result, err := service.foodRepository.DeleteFood(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service serviceFoods) GetFood(id int) (*Domain, error) {
	result, err := service.foodRepository.GetFood(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service serviceFoods) GetFoods() (*[]Domain, error) {
	result, err := service.foodRepository.GetFoods()
	if err != nil {
		return nil, err
	}
	return result, nil
}
