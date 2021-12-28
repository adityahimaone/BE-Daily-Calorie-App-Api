package nutritioninfo

import (
	"Daily-Calorie-App-API/businesses/foods"
	"Daily-Calorie-App-API/businesses/nutritioninfo"
	"gorm.io/gorm"
)

type NutritionInfo struct {
	gorm.Model
	ID          uint `gorm:"primary_key"`
	Calories    float64
	Fat         float64
	Carbs       float64
	Protein     float64
	ServingSize string
}

func fromDomain(domain nutritioninfo.Domain) NutritionInfo {
	return NutritionInfo{
		Calories:    domain.Calories,
		Fat:         domain.Fat,
		Carbs:       domain.Carbs,
		Protein:     domain.Protein,
		ServingSize: domain.ServingSize,
	}
}

func toDomain(domain NutritionInfo) (nutritioninfo.Domain, foods.Domain) {
	return nutritioninfo.Domain{
			ID:          int(domain.ID),
			Calories:    domain.Calories,
			Fat:         domain.Fat,
			Carbs:       domain.Carbs,
			Protein:     domain.Protein,
			ServingSize: domain.ServingSize,
		}, foods.Domain{
			NutritionInfoID: int(domain.ID),
			Calories:        domain.Calories,
			Fat:             domain.Fat,
			Carbs:           domain.Carbs,
			Protein:         domain.Protein,
			ServingSize:     domain.ServingSize,
		}
}
