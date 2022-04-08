package meal_plans_test

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/meal_plans"
	"Daily-Calorie-App-API/business/meal_plans/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockMealPlansRepository mocks.Repository
	mealPlanDomain          meal_plans.Domain
	mealPlanService         meal_plans.Service
)

func TestMain(m *testing.M) {
	mealPlanService = meal_plans.NewMealPlansService(&mockMealPlansRepository, &auth.ConfigJWT{})
	mealPlanDomain = meal_plans.Domain{
		ID:                 1,
		UserID:             1,
		DietaryPreferences: "balanced",
		PlanType:           "weekly",
		RangeCalories:      "2000-3000",
		MealPlans:          "test",
	}
	m.Run()
}

func TestMealPlansService_Create(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockMealPlansRepository.On("Insert", mock.Anything, mock.Anything).Return(&mealPlanDomain, nil).Once()
		inputMealPlan := meal_plans.Domain{
			UserID:             1,
			DietaryPreferences: "balanced",
			PlanType:           "weekly",
			RangeCalories:      "2000-3000",
			MealPlans:          "test",
		}
		result, err := mealPlanService.CreateMealPlans(&inputMealPlan)
		assert.Nil(t, err)
		assert.Equal(t, &mealPlanDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockMealPlansRepository.On("Insert", mock.Anything, mock.Anything).Return(&meal_plans.Domain{}, assert.AnError).Once()
		inputMealPlan := meal_plans.Domain{
			UserID:             1,
			DietaryPreferences: "balanced",
			PlanType:           "weekly",
			RangeCalories:      "2000-3000",
			MealPlans:          "test",
		}
		_, err := mealPlanService.CreateMealPlans(&inputMealPlan)
		assert.NotNil(t, err)
	})
}

func TestMealPlansService_GetLastMealPlans(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockMealPlansRepository.On("GetLastMealPlans", mock.Anything, mock.Anything).Return(&mealPlanDomain, nil).Once()
		result, err := mealPlanService.GetLastMealPlans(1)
		assert.Nil(t, err)
		assert.Equal(t, &mealPlanDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockMealPlansRepository.On("GetLastMealPlans", mock.Anything, mock.Anything).Return(&meal_plans.Domain{}, assert.AnError).Once()
		_, err := mealPlanService.GetLastMealPlans(1)
		assert.NotNil(t, err)
	})

}
