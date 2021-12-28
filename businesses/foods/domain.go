package foods

import (
	"Daily-Calorie-App-API/businesses/nutritioninfo"
	"time"
)

type Domain struct {
	ID              int
	Title           string
	ImgURL          string
	NutritionInfoID int
	Calories        float64
	Fat             float64
	Protein         float64
	Carbs           float64
	ServingSize     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Service interface {
	AddFood(food *Domain, nutritionInfo *nutritioninfo.Domain) (*Domain, error)
	EditFood(id int, food *Domain) (*Domain, error)
	DeleteFood(id int) (*Domain, error)
	GetFood(id int) (*Domain, error)
	GetFoods() (*[]Domain, error)
}

type Repository interface {
	InsertFood(food *Domain) (*Domain, error)
	UpdateFood(id int, food *Domain) (*Domain, error)
	DeleteFood(id int) (*Domain, error)
	GetFood(id int) (*Domain, error)
	GetFoods() (*[]Domain, error)
}
