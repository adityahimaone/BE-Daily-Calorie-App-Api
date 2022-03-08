package meal_plans

import (
	"time"
)

type Domain struct {
	ID                 int
	UserID             int
	DietaryPreferences string
	PlanType           string
	RangeCalories      string
	MealPlans          string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Service interface {
	CreateMealPlans(domain *Domain) (*Domain, error)
	GetLastMealPlans(userID int) (*Domain, error)
}

type Repository interface {
	Insert(domain *Domain) (*Domain, error)
	GetLastMealPlans(userID int) (*Domain, error)
}
