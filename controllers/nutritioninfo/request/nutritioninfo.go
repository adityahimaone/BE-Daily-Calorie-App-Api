package request

import (
	"Daily-Calorie-App-API/businesses/nutritioninfo"
	"gorm.io/gorm"
)

type NutritionInfo struct {
	gorm.Model
	ID          uint    `gorm:"primary_key"`
	Calories    float64 `json:"calories"`
	Carbs       float64 `json:"carbs"`
	Fat         float64 `json:"fat"`
	Protein     float64 `json:"protein"`
	ServingSize string  `json:"servingSize"`
}

func toDomain(request NutritionInfo) *nutritioninfo.Domain {
	return &nutritioninfo.Domain{
		Calories:    request.Calories,
		Carbs:       request.Carbs,
		Fat:         request.Fat,
		Protein:     request.Protein,
		ServingSize: request.ServingSize,
	}
}
