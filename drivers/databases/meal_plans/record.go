package meal_plans

import (
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
	DietaryPreferences string      `gorm:"type:varchar(100);"`
	PlanType           string      `gorm:"type:varchar(100);"`
	RangeCalories      string      `gorm:"type:varchar(100);"`
	MealPlans          string      `gorm:"type:json"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func fromDomain(domain meal_plans.Domain) MealPlans {
	return MealPlans{
		ID:                 uint(domain.ID),
		UserID:             uint(domain.UserID),
		DietaryPreferences: domain.DietaryPreferences,
		PlanType:           domain.PlanType,
		RangeCalories:      domain.RangeCalories,
		MealPlans:          domain.MealPlans,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}

func (record *MealPlans) toDomain() meal_plans.Domain {
	return meal_plans.Domain{
		ID:                 int(record.ID),
		UserID:             int(record.UserID),
		DietaryPreferences: record.DietaryPreferences,
		PlanType:           record.PlanType,
		RangeCalories:      record.RangeCalories,
		MealPlans:          record.MealPlans,
		CreatedAt:          record.CreatedAt,
		UpdatedAt:          record.UpdatedAt,
	}
}
