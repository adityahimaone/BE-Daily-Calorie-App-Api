package histories_detail_test

import (
	"Daily-Calorie-App-API/business/histories_detail"
	"Daily-Calorie-App-API/business/histories_detail/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockHistoriesDetailRepository mocks.Repository
	historiesDetailDomain         histories_detail.Domain
	historiesDetailService        histories_detail.Service
)

func TestMain(m *testing.M) {
	historiesDetailService = histories_detail.NewHistoriesDetailService(&mockHistoriesDetailRepository)
	historiesDetailDomain = histories_detail.Domain{
		ID:          1,
		FoodID:      1,
		HistoriesID: 1,
	}
	m.Run()
}

func TestHistoriesDetail_Insert(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("Insert", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		inputHistoriesDetail := histories_detail.Domain{
			ID:          1,
			FoodID:      1,
			HistoriesID: 1,
		}
		result, err := historiesDetailService.Insert(&inputHistoriesDetail)
		assert.Nil(t, err)
		assert.Equal(t, &historiesDetailDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("Insert", mock.Anything, mock.Anything).Return(&histories_detail.Domain{}, assert.AnError).Once()
		inputHistoriesDetail := histories_detail.Domain{
			ID:          1,
			FoodID:      1,
			HistoriesID: 1,
		}
		result, err := historiesDetailService.Insert(&inputHistoriesDetail)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistoriesDetail_Delete(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("Delete", mock.Anything, mock.Anything).Return(&historiesDetailDomain, nil).Once()
		result, err := historiesDetailService.Delete(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("Delete", mock.Anything, mock.Anything).Return(&histories_detail.Domain{}, assert.AnError).Once()
		result, err := historiesDetailService.Delete(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistoriesDetail_GetAllHistoriesDetailByHistoriesID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("GetAllHistoriesDetailByHistoriesID", mock.Anything, mock.Anything).Return(&[]histories_detail.Domain{historiesDetailDomain}, nil).Once()
		result, err := historiesDetailService.GetAllHistoriesDetailByHistoriesID(1)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("GetAllHistoriesDetailByHistoriesID", mock.Anything, mock.Anything).Return(&[]histories_detail.Domain{}, assert.AnError).Once()
		result, err := historiesDetailService.GetAllHistoriesDetailByHistoriesID(1)
		assert.NotNil(t, err)
		assert.NotNil(t, result)
	})
}

func TestHistoriesDetail_SumCalories(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("SumCalories", mock.Anything, mock.Anything).Return(float64(100), nil).Once()
		result, err := historiesDetailService.SumCalories(1)
		assert.Nil(t, err)
		assert.Equal(t, float64(100), result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockHistoriesDetailRepository.On("SumCalories", mock.Anything, mock.Anything).Return(float64(0), assert.AnError).Once()
		result, err := historiesDetailService.SumCalories(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, float64(100), result)
	})
}
