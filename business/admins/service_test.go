package admins_test

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business"
	"Daily-Calorie-App-API/business/admins"
	"Daily-Calorie-App-API/business/admins/mocks"
	"Daily-Calorie-App-API/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockAdminRepository mocks.Repository
	adminService        admins.Service
	adminDomain         admins.Domain
)

func TestMain(m *testing.M) {
	adminService = admins.NewServiceAdmin(&mockAdminRepository, &auth.ConfigJWT{})
	hashPassword, _ := helpers.PasswordHash("testMocking")
	adminDomain = admins.Domain{
		ID:       1,
		Username: "testMocking",
		Password: hashPassword,
	}
	m.Run()
}

func TestServiceAdmin_Login(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockAdminRepository.On("GetAdminByUsername", mock.Anything, mock.Anything).Return(&adminDomain, nil).Once()
		inputUser := admins.Domain{
			Username: "testMocking",
			Password: "testMocking",
		}
		result, err := adminService.Login(inputUser.Username, inputUser.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockAdminRepository.On("GetAdminByUsername", mock.Anything, mock.Anything).Return(&admins.Domain{}, assert.AnError).Once()
		inputUser := admins.Domain{
			Username: "testMocking",
			Password: "testMocking",
		}
		result, err := adminService.Login(inputUser.Username, inputUser.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockAdminRepository.On("GetAdminByUsername", mock.Anything, mock.Anything).Return(&adminDomain, nil).Once()
		inputUser := admins.Domain{
			Username: "testMocking",
			Password: "testMocking",
		}
		adminDomain.ID = 0
		result, err := adminService.Login(inputUser.Username, inputUser.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestServiceAdmin_Register(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockAdminRepository.On("Insert", mock.Anything, mock.Anything).Return(&adminDomain, nil).Once()
		inputUser := admins.Domain{
			Username: "testMocking",
			Password: "testMocking",
		}
		result, _ := adminService.Register(&inputUser)
		assert.Equal(t, &adminDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockAdminRepository.On("Insert", mock.Anything, mock.Anything).Return(&adminDomain, nil).Once()
		inputUser := admins.Domain{
			Username: "testMocking",
			Password: "testMocking",
		}
		result, err := adminService.Register(&inputUser)
		if result == &adminDomain {
			assert.Nil(t, err, business.ErrDuplicateData)
		}
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockAdminRepository.On("Insert", mock.Anything, mock.Anything).Return(&adminDomain, assert.AnError).Once()
		inputUser := admins.Domain{
			Username: "testMocking",
			Password: "testMocking",
		}
		_, err := adminService.Register(&inputUser)
		assert.NotNil(t, err)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockAdminRepository.On("Insert", mock.Anything, mock.Anything).Return(&adminDomain, assert.AnError).Once()
		inputUser := admins.Domain{
			Username: "testMocking",
			Password: "testMocking",
		}
		_, err := adminService.Register(&inputUser)
		assert.NotNil(t, err)
	})
}
