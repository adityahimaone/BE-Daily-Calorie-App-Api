package response

import (
	"Daily-Calorie-App-API/businesses/foods"
	"Daily-Calorie-App-API/businesses/nutritioninfo"
	"Daily-Calorie-App-API/controllers/nutritioninfo/response"
	"time"
)

type Food struct {
	Title           string                 `json:"title"`
	ImgURL          string                 `json:"img_url"`
	NutritionInfoID int                    `json:"nutrition_info_id"`
	NutritionInfo   response.NutritionInfo `json:"nutrition_info"`
	Calories        float64                `json:"calories"`
	Protein         float64                `json:"protein"`
	Carbs           float64                `json:"carbs"`
	Fat             float64                `json:"fat"`
	ServingSize     string                 `json:"serving_size"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updatedAt"`
}

func FromDomain(domain foods.Domain, domainNutrition nutritioninfo.Domain) Food {
	return Food{
		Title:           domain.Title,
		ImgURL:          domain.ImgURL,
		NutritionInfoID: domain.NutritionInfoID,
		Calories:        domainNutrition.Calories,
		Carbs:           domainNutrition.Carbs,
		Fat:             domainNutrition.Fat,
		Protein:         domainNutrition.Protein,
		ServingSize:     domainNutrition.ServingSize,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}

//func FromDomainArray(domain []foods.Domain, domainNutrition []nutritioninfo.Domain) []Food {
//	var response []Food
//	for i := 0; i < len(domain); i++ {
//		response = append(response, FromDomain(domain[i], domainNutrition[i]))
//	}
//	return response
//}
