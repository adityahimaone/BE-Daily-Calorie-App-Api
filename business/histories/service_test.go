package histories_test

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/foods"
	_mocksFoods "Daily-Calorie-App-API/business/foods/mocks"
	"Daily-Calorie-App-API/business/histories"
	"Daily-Calorie-App-API/business/histories/mocks"
	"Daily-Calorie-App-API/business/histories_detail"
	_mocksHistoriesDetail "Daily-Calorie-App-API/business/histories_detail/mocks"
	"Daily-Calorie-App-API/business/users"
	_mocksUsers "Daily-Calorie-App-API/business/users/mocks"
	"Daily-Calorie-App-API/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockHistoriesRepository    mocks.Repository
	mockHistoriesService       mocks.Service
	mockUsersService           _mocksUsers.Service
	mockFoodsService           _mocksFoods.Service
	mockHistoriesDetailService _mocksHistoriesDetail.Service
	usersDomain                users.Domain
	foodsDomain                foods.Domain
	historiesDetailDomain      histories_detail.Domain
	historiesDomain            histories.Domain
	hitoriesService            histories.Service
)

func TestMain(m *testing.M) {
	hitoriesService = histories.NewHistoriesService(&mockHistoriesRepository, &mockUsersService, &mockFoodsService, &mockHistoriesDetailService, &auth.ConfigJWT{})
	hashPassword, _ := helpers.PasswordHash("test")
	historiesDomain = histories.Domain{
		ID:            1,
		UserID:        1,
		FoodID:        1,
		Water:         1,
		TotalCalories: 1000,
		TotalFood:     4,
		Date:          "02-01-2006",
	}
	usersDomain = users.Domain{
		ID:             1,
		Name:           "test",
		Password:       hashPassword,
		Email:          "test@mail.com",
		Gender:         "male",
		PersonalDataID: 1,
	}
	foodsDomain = foods.Domain{
		ID:       1,
		Title:    "test",
		Calories: 1000,
		Carbs:    10,
		Protein:  10,
		Fat:      10,
	}
	historiesDetailDomain = histories_detail.Domain{
		ID:          1,
		FoodID:      1,
		HistoriesID: 1,
	}
	m.Run()
}

func TestHistories_CreateHistoriesFromAPI(t *testing.T) {
	t.Run("Valid Test CreateHistoriesFromAPI Food Exist", func(t *testing.T) {
		mockFoodsService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		inputFood := foods.Domain{
			ID:       1,
			Title:    "test",
			Calories: 1000,
			Carbs:    10,
			Protein:  10,
			Fat:      10,
		}
		mockHistoriesService.On("CreateHistories", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		result, err := hitoriesService.CreateHistoriesFromAPI(&inputHistories, &inputFood)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Valid Test CreateHistoriesFromAPI Food Not Exist", func(t *testing.T) {
		mockFoodsService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		inputFood := foods.Domain{
			ID:       1,
			Title:    "test",
			Calories: 1000,
			Carbs:    10,
			Protein:  10,
			Fat:      10,
		}
		foodsDomain.ID = 0
		if foodsDomain.ID == 0 {
			mockFoodsService.On("AddFood", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
			mockHistoriesService.On("CreateHistories", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		}
		result, err := hitoriesService.CreateHistoriesFromAPI(&inputHistories, &inputFood)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test CreateHistoriesFromAPI GetFoodByName", func(t *testing.T) {
		mockFoodsService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		inputFood := foods.Domain{
			ID:       1,
			Title:    "test",
			Calories: 1000,
			Carbs:    10,
			Protein:  10,
			Fat:      10,
		}
		result, err := hitoriesService.CreateHistoriesFromAPI(&inputHistories, &inputFood)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test CreateHistoriesFromAPI Food Not Exist Add Food", func(t *testing.T) {
		mockFoodsService.On("GetFoodByName", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		inputFood := foods.Domain{
			ID:       1,
			Title:    "test",
			Calories: 1000,
			Carbs:    10,
			Protein:  10,
			Fat:      10,
		}
		foodsDomain.ID = 0
		if foodsDomain.ID == 0 {
			mockFoodsService.On("AddFood", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
			//mockHistoriesService.On("CreateHistories", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		}
		result, err := hitoriesService.CreateHistoriesFromAPI(&inputHistories, &inputFood)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistories_CreateWater(t *testing.T) {
	t.Run("Valid Test Water Exist", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("UpdateWater", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		}
		result, err := hitoriesService.CreateWater(&inputHistories)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Valid Test Water Not Exist", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		historiesDomain.ID = 0
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("UpdateWater", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		} else {
			mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		}
		result, err := hitoriesService.CreateWater(&inputHistories)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid GetUserByID", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&users.Domain{}, assert.AnError).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		result, err := hitoriesService.CreateWater(&inputHistories)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test Water Exist", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("UpdateWater", mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		}
		result, err := hitoriesService.CreateWater(&inputHistories)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("invalid Test Water Not Exist", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		historiesDomain.ID = 0
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("UpdateWater", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		} else {
			mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		}
		result, err := hitoriesService.CreateWater(&inputHistories)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistories_CreateHistories(t *testing.T) {
	t.Run("Valid Test Histories Exist", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
			mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Valid Test Histories Not Exist", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		historiesDomain.ID = 0
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
			mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		} else {
			mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
			mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid GetUserByID", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&users.Domain{}, assert.AnError).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid GetFoodByID", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foods.Domain{}, assert.AnError).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Histories Exist Update", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Histories Exist Insert", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
			mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&histories_detail.Domain{}, assert.AnError).Once()
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
	t.Run("Invalid Test Histories Not Exist Insert Histories", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		historiesDomain.ID = 0
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
			mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		} else {
			mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test Histories Not Exist Insert Histories Detail", func(t *testing.T) {
		mockUsersService.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockFoodsService.On("GetFoodByID", mock.Anything, mock.Anything).Return(&foodsDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		inputHistories := histories.Domain{
			UserID: 1,
			FoodID: 1,
			Water:  1,
			Date:   "02-01-2006",
		}
		historiesDomain.ID = 0
		if historiesDomain.ID != 0 {
			mockHistoriesRepository.On("Update", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
			mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		} else {
			mockHistoriesRepository.On("Insert", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
			mockHistoriesDetailService.On("Insert", mock.Anything, mock.Anything).Return(&histories_detail.Domain{}, assert.AnError).Once()
		}
		result, err := hitoriesService.CreateHistories(&inputHistories)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}

func TestHistories_GetLastHistoryByUserID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetLastHistoryByUserID", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		result, err := hitoriesService.GetLastHistoryByUserID(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetLastHistoryByUserID", mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.GetLastHistoryByUserID(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistories_GetAllHistoryByUserID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetAllHistoryByUserID", mock.Anything, mock.Anything).Return(&[]histories.Domain{historiesDomain}, nil).Once()
		result, err := hitoriesService.GetAllHistoryByUserID(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetAllHistoryByUserID", mock.Anything, mock.Anything).Return(&[]histories.Domain{histories.Domain{}}, assert.AnError).Once()
		result, err := hitoriesService.GetAllHistoryByUserID(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistories_GetHistoriesByUserIDandDate(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		result, err := hitoriesService.GetHistoriesByUserIDandDate(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetHistoriesByUserIDandDate", mock.Anything, mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.GetHistoriesByUserIDandDate(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistories_GetHistoriesByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetHistoriesByID", mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		result, err := hitoriesService.GetHistoriesByID(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesRepository.On("GetHistoriesByID", mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.GetHistoriesByID(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistories_DeleteHistoriesDetail(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesDetailService.On("Delete", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		mockHistoriesDetailService.On("SumCalories", mock.Anything, mock.Anything).Return(100.0, nil).Once()
		mockHistoriesRepository.On("UpdateTotalCalories", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByID", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("UpdateTotalFood", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		result, err := hitoriesService.DeleteHistoriesDetail(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test Delete", func(t *testing.T) {
		mockHistoriesDetailService.On("Delete", mock.Anything, mock.Anything).Return(&histories_detail.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.DeleteHistoriesDetail(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test SumCalories", func(t *testing.T) {
		mockHistoriesDetailService.On("Delete", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		mockHistoriesDetailService.On("SumCalories", mock.Anything, mock.Anything).Return(100.0, assert.AnError).Once()
		result, err := hitoriesService.DeleteHistoriesDetail(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test UpdateTotalCalories", func(t *testing.T) {
		mockHistoriesDetailService.On("Delete", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		mockHistoriesDetailService.On("SumCalories", mock.Anything, mock.Anything).Return(100.0, nil).Once()
		mockHistoriesRepository.On("UpdateTotalCalories", mock.Anything, mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.DeleteHistoriesDetail(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test GetHistoriesByID", func(t *testing.T) {
		mockHistoriesDetailService.On("Delete", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		mockHistoriesDetailService.On("SumCalories", mock.Anything, mock.Anything).Return(100.0, nil).Once()
		mockHistoriesRepository.On("UpdateTotalCalories", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByID", mock.Anything, mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.DeleteHistoriesDetail(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test UpdateTotalFood", func(t *testing.T) {
		mockHistoriesDetailService.On("Delete", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		mockHistoriesDetailService.On("SumCalories", mock.Anything, mock.Anything).Return(100.0, nil).Once()
		mockHistoriesRepository.On("UpdateTotalCalories", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("GetHistoriesByID", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		mockHistoriesRepository.On("UpdateTotalFood", mock.Anything, mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.DeleteHistoriesDetail(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistories_UpdateTotalCalories(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesRepository.On("UpdateTotalCalories", mock.Anything, mock.Anything, mock.Anything).Return(&historiesDomain, nil).Once()
		result, err := hitoriesService.UpdateTotalCalories(1, 1000)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesRepository.On("UpdateTotalCalories", mock.Anything, mock.Anything, mock.Anything).Return(&histories.Domain{}, assert.AnError).Once()
		result, err := hitoriesService.UpdateTotalCalories(1, 1000)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}
