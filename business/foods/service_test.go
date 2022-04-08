package foods_test

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/foods"
	"Daily-Calorie-App-API/business/foods/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockFoodRepository mocks.Repository
	foodsDomain        foods.Domain
	foodsService       foods.Service
)

func TestMain(m *testing.M) {
	foodsService = foods.NewFoodService(&mockFoodRepository, &auth.ConfigJWT{})
	foodsDomain = foods.Domain{
		ID:          1,
		Title:       "food test",
		ImgURL:      "test.img",
		Calories:    100,
		Fat:         10,
		Carbs:       20,
		Protein:     30,
		ServingSize: 100,
	}
	m.Run()
}
func TestFoodService_AddFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		inputFood := foods.Domain{
			Title:       "food test",
			ImgURL:      "test.img",
			Calories:    100,
			Fat:         10,
			Carbs:       20,
			Protein:     30,
			ServingSize: 100,
		}
		result, err := foodsService.AddFood(&inputFood)
		assert.Nil(t, err)
		assert.Equal(t, &foodsDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Insert", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		inputFood := foods.Domain{
			Title:       "food test",
			ImgURL:      "test.img",
			Calories:    100,
			Fat:         10,
			Carbs:       20,
			Protein:     30,
			ServingSize: 100,
		}
		result, err := foodsService.AddFood(&inputFood)
		assert.NotNil(t, err)
		assert.Equal(t, &foods.Domain{}, result)
	})
}

func TestFoodService_EditFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("Update", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		inputFood := foods.Domain{
			ID:          1,
			Title:       "food test",
			ImgURL:      "test.img",
			Calories:    100,
			Fat:         10,
			Carbs:       20,
			Protein:     30,
			ServingSize: 100,
		}
		result, err := foodsService.EditFood(1, &inputFood)
		assert.Nil(t, err)
		assert.Equal(t, &foodsDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Update", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		inputFood := foods.Domain{
			ID:          1,
			Title:       "food test",
			ImgURL:      "test.img",
			Calories:    100,
			Fat:         10,
			Carbs:       20,
			Protein:     30,
			ServingSize: 100,
		}
		result, err := foodsService.EditFood(1, &inputFood)
		assert.NotNil(t, err)
		assert.Equal(t, &foods.Domain{}, result)
	})
}

func TestFoodService_DeleteFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("Delete", mock.Anything, mock.Anything).Return("Delete Food Successfully", nil).Once()
		result, err := foodsService.DeleteFood(1)
		assert.Nil(t, err)
		assert.Equal(t, "Delete Food Successfully", result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("Delete", mock.Anything, mock.Anything).Return("", assert.AnError).Once()
		result, err := foodsService.DeleteFood(1)
		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})
}

func TestFoodService_GetFoodByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		result, err := foodsService.GetFoodByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &foodsDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		result, err := foodsService.GetFoodByID(1)
		assert.NotNil(t, err)
		assert.Equal(t, &foods.Domain{}, result)
	})
}

func TestFoodService_GetFoodByName(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		result, err := foodsService.GetFoodByName("food test")
		assert.Nil(t, err)
		assert.Equal(t, &foodsDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		result, err := foodsService.GetFoodByName("food test")
		assert.NotNil(t, err)
		assert.Equal(t, &foods.Domain{}, result)
	})
}

func TestFoodService_GetAllFood(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockFoodRepository.On("GetAllFood", mock.Anything, mock.Anything).Return(&[]foods.Domain{foodsDomain}, nil).Once()
		result, err := foodsService.GetAllFood()
		assert.Nil(t, err)
		assert.Equal(t, &[]foods.Domain{foodsDomain}, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockFoodRepository.On("GetAllFood", mock.Anything, mock.Anything).Return(&[]foods.Domain{}, assert.AnError).Once()
		result, err := foodsService.GetAllFood()
		assert.NotNil(t, err)
		assert.Equal(t, &[]foods.Domain{}, result)
	})
}
