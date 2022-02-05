package request

import (
	"Daily-Calorie-App-API/business/foods"
)

type Food struct {
	Title    string  `json:"title"`
	ImgURL   string  `json:"img_url"`
	Calories float64 `json:"calories"`
	Carbs    float64 `json:"carbs"`
	Protein  float64 `json:"protein"`
	Fat      float64 `json:"fat"`
}

func (request *Food) ToDomain() *foods.Domain {
	return &foods.Domain{
		Title:    request.Title,
		ImgURL:   request.ImgURL,
		Calories: request.Calories,
		Carbs:    request.Carbs,
		Fat:      request.Fat,
		Protein:  request.Protein,
	}
}
