package meal_plans

import (
	"Daily-Calorie-App-API/business/foodsAPI"
	"Daily-Calorie-App-API/business/meal_plans"
	"Daily-Calorie-App-API/drivers/databases/users"
	"gorm.io/gorm"
	"time"
)

type MealPlans struct {
	gorm.Model
	ID                 uint        `gorm:"primary_key"`
	UserID             uint        `gorm:"not null"`
	User               users.Users `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
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

func fromDomain(domain meal_plans.Domain) MealPlans {
	return MealPlans{
		ID:                 uint(domain.ID),
		UserID:             uint(domain.UserID),
		MealTime:           domain.MealTime,
		DietaryPreferences: domain.DietaryPreferences,
		PlanType:           domain.PlanType,
		RangeCalories:      domain.RangeCalories,
		Lunch:              domain.Lunch,
		Dinner:             domain.Dinner,
		Breakfast:          domain.Breakfast,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}

func (record *MealPlans) toDomain() meal_plans.Domain {
	return meal_plans.Domain{
		ID:                 int(record.ID),
		UserID:             int(record.UserID),
		MealTime:           record.MealTime,
		DietaryPreferences: record.DietaryPreferences,
		PlanType:           record.PlanType,
		RangeCalories:      record.RangeCalories,
		Lunch:              record.Lunch,
		Dinner:             record.Dinner,
		Breakfast:          record.Breakfast,
		CreatedAt:          record.CreatedAt,
		UpdatedAt:          record.UpdatedAt,
	}
}
