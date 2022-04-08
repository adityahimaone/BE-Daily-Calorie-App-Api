package response

import (
	"Daily-Calorie-App-API/business/histories_detail"
	"time"
)

type HistoriesDetail struct {
	ID           int       `json:"id"`
	HistoriesID  int       `json:"histories_id"`
	FoodID       int       `json:"food_id"`
	FoodTitle    string    `json:"food_name"`
	FoodImage    string    `json:"food_image"`
	FoodCalories float64   `json:"food_calories"`
	FoodCarbs    float64   `json:"food_carbs"`
	FoodProtein  float64   `json:"food_protein"`
	FoodFat      float64   `json:"food_fat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain histories_detail.Domain) HistoriesDetail {
	return HistoriesDetail{
		ID:           domain.ID,
		HistoriesID:  domain.HistoriesID,
		FoodID:       domain.FoodID,
		FoodTitle:    domain.FoodTitle,
		FoodImage:    domain.FoodImage,
		FoodCalories: domain.FoodCalories,
		FoodCarbs:    domain.FoodCrabs,
		FoodProtein:  domain.FoodProtein,
		FoodFat:      domain.FoodFat,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromDomainArray(domains []histories_detail.Domain) []HistoriesDetail {
	var historiesDetails []HistoriesDetail
	for _, domain := range domains {
		historiesDetails = append(historiesDetails, FromDomain(domain))
	}
	return historiesDetails
}
