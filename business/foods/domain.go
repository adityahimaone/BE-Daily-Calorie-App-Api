package foods

import (
	"time"
)

type Domain struct {
	ID          int
	Title       string
	ImgURL      string
	Calories    float64
	Fat         float64
	Protein     float64
	Carbs       float64
	ServingSize string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	AddFood(food *Domain) (*Domain, error)
	EditFood(id int, food *Domain) (*Domain, error)
	DeleteFood(id int) (string, error)
	GetFoodByID(id int) (*Domain, error)
	GetAllFood() (*[]Domain, error)
	GetFoodByName(name string) (*Domain, error)
}

type Repository interface {
	Insert(food *Domain) (*Domain, error)
	Update(id int, food *Domain) (*Domain, error)
	Delete(id int) (string, error)
	GetFoodByID(id int) (*Domain, error)
	GetAllFood() (*[]Domain, error)
	GetFoodByName(name string) (*Domain, error)
}
