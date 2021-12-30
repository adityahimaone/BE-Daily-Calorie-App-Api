package request

import (
	"Daily-Calorie-App-API/businesses/foods"
	"Daily-Calorie-App-API/businesses/nutritioninfo"
	"Daily-Calorie-App-API/controllers/nutritioninfo/request"
)

type Food struct {
	Title           string                `json:"title"`
	ImgURL          string                `json:"img_url"`
	NutritionInfoID int                   `json:"nutrition_info_id"`
	NutritionInfo   request.NutritionInfo `json:"nutrition_info"`
}

func ToDomain(request Food) (*foods.Domain, *nutritioninfo.Domain) {
	return &foods.Domain{
			Title:           request.Title,
			ImgURL:          request.ImgURL,
			NutritionInfoID: request.NutritionInfoID,
		}, &nutritioninfo.Domain{
			Calories:    request.NutritionInfo.Calories,
			Carbs:       request.NutritionInfo.Carbs,
			Fat:         request.NutritionInfo.Fat,
			Protein:     request.NutritionInfo.Protein,
			ServingSize: request.NutritionInfo.ServingSize,
		}
}
