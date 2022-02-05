package request

import "Daily-Calorie-App-API/business/histories"

type Histories struct {
	UserID int `json:"user_id"`
	Water  int `json:"water"`
	FoodID int `json:"food_id"`
}

func (request *Histories) ToDomain() *histories.Domain {
	return &histories.Domain{
		UserID: request.UserID,
		Water:  request.Water,
		FoodID: request.FoodID,
	}
}
