package response

import (
	"Daily-Calorie-App-API/businesses/foods"
	"time"
)

type Food struct {
	Title           string    `json:"title"`
	ImgURL          string    `json:"img_url"`
	NutritionInfoID int       `json:"nutrition_info_id"`
	Calories        float64   `json:"calories"`
	Protein         float64   `json:"protein"`
	Carbs           float64   `json:"carbs"`
	Fat             float64   `json:"fat"`
	ServingSize     string    `json:"serving_size"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func FromDomain(domain foods.Domain) Food {
	return Food{
		Title:           domain.Title,
		ImgURL:          domain.ImgURL,
		NutritionInfoID: domain.NutritionInfoID,
		Calories:        domain.Calories,
		Carbs:           domain.Carbs,
		Fat:             domain.Fat,
		Protein:         domain.Protein,
		ServingSize:     domain.ServingSize,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}

func FromDomainArray(domains []foods.Domain) []Food {
	var foods []Food
	for _, domain := range domains {
		foods = append(foods, FromDomain(domain))
	}
	return foods
}
