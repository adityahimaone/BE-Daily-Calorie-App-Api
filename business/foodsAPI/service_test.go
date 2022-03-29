package foodsAPI_test

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/foodsAPI"
	"Daily-Calorie-App-API/business/foodsAPI/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockFoodAPIRepository mocks.Repository
	foodAPIDomain         foodsAPI.Domain
	recipeDomain          foodsAPI.DomainRecipe
	foodAPIService        foodsAPI.Service
)

func TestMain(m *testing.M) {
	foodAPIService = foodsAPI.NewFoodAPIService(&mockFoodAPIRepository, &auth.ConfigJWT{})
	foodAPIDomain = foodsAPI.Domain{
		ID:             1,
		Title:          "nasi goreng",
		ImgURL:         "test/png",
		Calories:       100,
		Fat:            10,
		Protein:        10,
		Carbs:          10,
		ServingSize:    10,
		RecipeLabel:    "test",
		RecipeImageURL: "test/png",
		RecipeURL:      "test",
		RecipeCalories: 100,
		RecipeIngredients: []string{
			"test",
		},
	}
	recipeDomain = foodsAPI.DomainRecipe{
		RecipeLabel:    "test",
		RecipeImageURL: "test/png",
		RecipeURL:      "test",
		RecipeCalories: 100,
		RecipeIngredients: []string{
			"test",
		},
	}
	m.Run()
}

func TestFoodAPI_GetFoodByName(t *testing.T) {
	t.Run("Valid Test GetFoodByName", func(t *testing.T) {
		mockFoodAPIRepository.On("GetFoodByName", mock.Anything, mock.Anything).Return(&[]foodsAPI.Domain{foodAPIDomain}, nil).Once()
		result, err := foodAPIService.GetFoodByName("nasi goreng")
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test GetFoodByName", func(t *testing.T) {
		mockFoodAPIRepository.On("GetFoodByName", mock.Anything, mock.Anything).Return(&[]foodsAPI.Domain{}, assert.AnError).Once()
		result, err := foodAPIService.GetFoodByName("nasi goreng")
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestFoodAPI_GetMealPlan(t *testing.T) {
	t.Run("Valid Test GetMealPlan", func(t *testing.T) {
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{recipeDomain}, nil).Once()
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{recipeDomain}, nil).Once()
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{recipeDomain}, nil).Once()
		inputMealPlan := foodsAPI.Domain{
			MealTime:           "lunch",
			DietaryPreferences: "balanced",
			PlanType:           "weekly",
			RangeCalories:      "100-200",
		}
		result, err := foodAPIService.GetMealPlan(&inputMealPlan)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test GetMealPlan Breakfast", func(t *testing.T) {
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{recipeDomain}, nil).Once()
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{}, assert.AnError).Once()
		inputMealPlan := foodsAPI.Domain{
			MealTime:           "lunch",
			DietaryPreferences: "balanced",
			PlanType:           "weekly",
			RangeCalories:      "100-200",
		}
		result, err := foodAPIService.GetMealPlan(&inputMealPlan)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test GetMealPlan Lunch", func(t *testing.T) {
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{}, assert.AnError).Once()
		inputMealPlan := foodsAPI.Domain{
			MealTime:           "lunch",
			DietaryPreferences: "balanced",
			PlanType:           "weekly",
			RangeCalories:      "100-200",
		}
		result, err := foodAPIService.GetMealPlan(&inputMealPlan)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test GetMealPlan Dinner", func(t *testing.T) {
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{recipeDomain}, nil).Once()
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{recipeDomain}, nil).Once()
		mockFoodAPIRepository.On("GetMealPlan", mock.Anything, mock.Anything).Return(&[]foodsAPI.DomainRecipe{}, assert.AnError).Once()
		inputMealPlan := foodsAPI.Domain{
			MealTime:           "lunch",
			DietaryPreferences: "balanced",
			PlanType:           "weekly",
			RangeCalories:      "100-200",
		}
		result, err := foodAPIService.GetMealPlan(&inputMealPlan)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}
