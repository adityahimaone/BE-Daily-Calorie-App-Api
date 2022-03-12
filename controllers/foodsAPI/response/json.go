package response

import (
	"Daily-Calorie-App-API/business/foodsAPI"
)

type Food struct {
	Title       string  `json:"title"`
	ImgURL      string  `json:"img_url"`
	Calories    float64 `json:"calories"`
	Protein     float64 `json:"protein"`
	Carbs       float64 `json:"carbs"`
	Fat         float64 `json:"fat"`
	ServingSize float64 `json:"serving_size"`
}

type MealPlan struct {
	Lunch     []foodsAPI.DomainRecipe `json:"lunch"`
	Dinner    []foodsAPI.DomainRecipe `json:"dinner"`
	Breakfast []foodsAPI.DomainRecipe `json:"breakfast"`
}

func FromDomain(domain foodsAPI.Domain) Food {
	return Food{
		Title:       domain.Title,
		ImgURL:      domain.ImgURL,
		Calories:    domain.Calories,
		Carbs:       domain.Carbs,
		Fat:         domain.Fat,
		Protein:     domain.Protein,
		ServingSize: domain.ServingSize,
	}
}

func FromDomainRecipe(domain foodsAPI.Domain) MealPlan {
	return MealPlan{
		Lunch:     domain.Lunch,
		Dinner:    domain.Dinner,
		Breakfast: domain.Breakfast,
	}
}

func FromDomainArray(domains []foodsAPI.Domain) []Food {
	var foodsArray []Food
	for _, domain := range domains {
		foodsArray = append(foodsArray, FromDomain(domain))
	}
	return foodsArray
}
