package meal_plans

import (
	"Daily-Calorie-App-API/business/foodsAPI"
	"time"
)

type Domain struct {
	ID                 int
	UserID             int
	MealTime           string
	DietaryPreferences string
	PlanType           string
	RangeCalories      string
	Lunch              []foodsAPI.DomainRecipe
	Dinner             []foodsAPI.DomainRecipe
	Breakfast          []foodsAPI.DomainRecipe
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Service interface {
	CreateMealPlans(domain *Domain, domainFoodAPI *foodsAPI.Domain) (*Domain, error)
}

type Repository interface {
	Insert(domain *Domain) (*Domain, error)
}
