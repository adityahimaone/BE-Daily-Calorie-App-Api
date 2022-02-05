package response

import (
	"Daily-Calorie-App-API/business/foods"
	"time"
)

type Food struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	ImgURL      string    `json:"img_url"`
	Calories    float64   `json:"calories"`
	Protein     float64   `json:"protein"`
	Carbs       float64   `json:"carbs"`
	Fat         float64   `json:"fat"`
	ServingSize string    `json:"serving_size"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain foods.Domain) Food {
	return Food{
		ID:          int64(domain.ID),
		Title:       domain.Title,
		ImgURL:      domain.ImgURL,
		Calories:    domain.Calories,
		Carbs:       domain.Carbs,
		Fat:         domain.Fat,
		Protein:     domain.Protein,
		ServingSize: domain.ServingSize,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromDomainArray(domains []foods.Domain) []Food {
	var foodsArray []Food
	for _, domain := range domains {
		foodsArray = append(foodsArray, FromDomain(domain))
	}
	return foodsArray
}
