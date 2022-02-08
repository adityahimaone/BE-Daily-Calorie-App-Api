package request

import (
	"Daily-Calorie-App-API/business/foods"
	"Daily-Calorie-App-API/business/histories"
)

type Histories struct {
	UserID int `json:"user_id"`
	Water  int `json:"water"`
	FoodID int `json:"food_id"`
}

type HistoriesFood struct {
	UserID      int     `json:"user_id"`
	FoodID      int     `json:"food_id"`
	Title       string  `json:"title"`
	ImgURL      string  `json:"img_url"`
	Calories    float64 `json:"calories"`
	Carbs       float64 `json:"carbs"`
	Protein     float64 `json:"protein"`
	Fat         float64 `json:"fat"`
	ServingSize float64 `json:"serving_size"`
}

func (request *Histories) ToDomain() *histories.Domain {
	return &histories.Domain{
		UserID: request.UserID,
		Water:  request.Water,
		FoodID: request.FoodID,
	}
}

func ToDomain(request *HistoriesFood) (*foods.Domain, *histories.Domain) {
	return &foods.Domain{
			Title:       request.Title,
			ImgURL:      request.ImgURL,
			Calories:    request.Calories,
			Carbs:       request.Carbs,
			Protein:     request.Protein,
			Fat:         request.Fat,
			ServingSize: request.ServingSize,
		}, &histories.Domain{
			UserID: request.UserID,
			FoodID: request.FoodID,
		}
}
