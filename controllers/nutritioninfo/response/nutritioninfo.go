package response

type NutritionInfo struct {
	Calories    float64 `json:"calories"`
	Carbs       float64 `json:"carbs"`
	Fat         float64 `json:"fat"`
	Protein     float64 `json:"protein"`
	ServingSize string  `json:"servingSize"`
}

func FromDomain(nutritionInfo NutritionInfo) NutritionInfo {
	return NutritionInfo{
		Calories:    nutritionInfo.Calories,
		Carbs:       nutritionInfo.Carbs,
		Fat:         nutritionInfo.Fat,
		Protein:     nutritionInfo.Protein,
		ServingSize: nutritionInfo.ServingSize,
	}
}
