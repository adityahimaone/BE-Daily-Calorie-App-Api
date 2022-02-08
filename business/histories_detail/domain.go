package histories_detail

import "time"

type Domain struct {
	ID           int `json:"id"`
	FoodID       int `json:"food_id"`
	FoodTitle    string
	FoodImage    string
	FoodCalories float64
	FoodCrabs    float64
	FoodProtein  float64
	FoodFat      float64
	HistoriesID  int       `json:"histories_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Service interface {
	Insert(historiesDetail *Domain) (*Domain, error)
	GetAllHistoriesDetailByHistoriesID(historiesID int) (*[]Domain, error)
	Delete(id int) (string, error)
}

type Repository interface {
	Insert(historiesDetail *Domain) (*Domain, error)
	GetAllHistoriesDetailByHistoriesID(historiesID int) (*[]Domain, error)
	Delete(id int) error
}
