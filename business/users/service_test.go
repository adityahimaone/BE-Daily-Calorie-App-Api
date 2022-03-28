package users_test

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/personal_data"
	_personalDataMocks "Daily-Calorie-App-API/business/personal_data/mocks"
	"Daily-Calorie-App-API/business/users"
	"Daily-Calorie-App-API/business/users/mocks"
	"Daily-Calorie-App-API/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mockUserRepository mocks.Repository
	//mockUserService            mocks.Service
	mockPersonalDataRepository _personalDataMocks.Repository
	usersService               users.Service
	usersDomain                users.Domain
	personalDataDomain         personal_data.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserService(&mockUserRepository, &mockPersonalDataRepository, &auth.ConfigJWT{})
	hashPassword, _ := helpers.PasswordHash("testUser")
	usersDomain = users.Domain{
		ID:             1,
		Email:          "testuser@mail.com",
		Password:       hashPassword,
		Name:           "Test User",
		AvatarUrl:      "ava.png",
		Gender:         "male",
		PersonalDataID: 1,
		Calories:       1000,
		Weight:         50,
		Height:         160,
	}
	personalDataDomain = personal_data.Domain{
		ID:       1,
		Calories: 1000,
		Weight:   50,
		Height:   160,
	}
	m.Run()
}

func TestLogin(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		inputUser := users.Domain{
			Email:    "testuser@mail.com",
			Password: "testUser",
		}
		result, err := usersService.Login(inputUser.Email, inputUser.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&users.Domain{}, assert.AnError).Once()
		inputUser := users.Domain{
			Email:    "testuser@mail.com",
			Password: "testUser",
		}
		result, err := usersService.Login(inputUser.Email, inputUser.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		inputUser := users.Domain{
			Email:    "testuser@mail.com",
			Password: "testUser",
		}
		usersDomain.ID = 0
		result, err := usersService.Login(inputUser.Email, inputUser.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestCountCalories(t *testing.T) {
	t.Run("Valid Test Male", func(t *testing.T) {
		inputUser := users.Domain{
			Gender:       "male",
			Age:          25,
			Weight:       50,
			Height:       160,
			ActivityType: 1.5,
		}
		inputPersonalData := personal_data.Domain{
			Weight: 50,
			Height: 160,
		}
		result, err := usersService.CountCalories(&inputUser, &inputPersonalData)
		expectedResult := 1312.5
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
	t.Run("Valid Female", func(t *testing.T) {
		inputUser := users.Domain{
			Gender:       "female",
			Age:          25,
			Weight:       50,
			Height:       160,
			ActivityType: 1.5,
		}
		inputPersonalData := personal_data.Domain{
			Weight: 50,
			Height: 160,
		}
		result, err := usersService.CountCalories(&inputUser, &inputPersonalData)
		expectedResult := 1821.0
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, result)
	})
}

func TestServiceUsers_RegisterUser(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		//expectedResult := 1312.5
		//mockUserService.On("CountCalories", mock.Anything, mock.Anything).Return(expectedResult, nil).Once()
		mockPersonalDataRepository.On("Insert", mock.Anything, mock.Anything).Return(&personalDataDomain, nil).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		inputUser := users.Domain{
			Email:          "testuser@mail.com",
			Password:       "testUser",
			AvatarUrl:      "ava.png",
			Gender:         "male",
			PersonalDataID: 1,
			Age:            25,
			Weight:         50,
			Height:         160,
			ActivityType:   1.5,
		}
		inputPersonalData := personal_data.Domain{
			ID:       1,
			Calories: 1000,
			Weight:   50,
			Height:   160,
		}
		result, err := usersService.RegisterUser(&inputUser, &inputPersonalData)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		//mockUserService.On("CountCalories", mock.Anything, mock.Anything).Return(nil, assert.AnError).Once()
		mockPersonalDataRepository.On("Insert", mock.Anything, mock.Anything).Return(&personalDataDomain, nil).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		_, err := usersService.RegisterUser(&users.Domain{}, &personal_data.Domain{})
		assert.Nil(t, err)
	})
	t.Run("Invalid Test Personal Data Insert", func(t *testing.T) {
		//expectedResult := 1312.5
		//mockUserService.On("CountCalories", mock.Anything, mock.Anything).Return(expectedResult, nil).Once()
		mockPersonalDataRepository.On("Insert", mock.Anything, mock.Anything).Return(&personal_data.Domain{}, assert.AnError).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		inputUser := users.Domain{
			Email:          "testuser@mail.com",
			Password:       "testUser",
			AvatarUrl:      "ava.png",
			Gender:         "male",
			PersonalDataID: 1,
			Age:            25,
			Weight:         50,
			Height:         160,
			ActivityType:   5,
		}
		inputPersonalData := personal_data.Domain{
			ID:       1,
			Calories: 1000,
			Weight:   50,
			Height:   160,
		}
		result, err := usersService.RegisterUser(&inputUser, &inputPersonalData)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
	t.Run("Invalid Test Users Insert", func(t *testing.T) {
		//expectedResult := 1312.5
		//mockUserService.On("CountCalories", mock.Anything, mock.Anything).Return(expectedResult, nil).Once()
		mockPersonalDataRepository.On("Insert", mock.Anything, mock.Anything).Return(&personalDataDomain, nil).Once()
		mockUserRepository.On("Insert", mock.Anything, mock.Anything).Return(&users.Domain{}, assert.AnError).Once()
		inputUser := users.Domain{
			Email:          "testuser@mail.com",
			Password:       "testUser",
			AvatarUrl:      "ava.png",
			Gender:         "male",
			PersonalDataID: 1,
			Age:            25,
			Weight:         50,
			Height:         160,
			ActivityType:   5,
		}
		inputPersonalData := personal_data.Domain{
			ID:       1,
			Calories: 1000,
			Weight:   50,
			Height:   160,
		}
		result, err := usersService.RegisterUser(&inputUser, &inputPersonalData)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		result, err := usersService.GetUserByID(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("GetUserByID", mock.Anything, mock.Anything).Return(&usersDomain, assert.AnError).Once()
		result, err := usersService.GetUserByID(1)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestEmailAvailable(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		result, _ := usersService.EmailAvailable("testuser@mail.com")
		if result {
			assert.False(t, false)
		}
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("GetUserByEmail", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		result, _ := usersService.EmailAvailable("testuser@mail.com")
		userID := 0
		if &usersDomain.ID == &userID {
			assert.True(t, result)
		}
	})
}

func TestGetAllUser(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("GetAllUser", mock.Anything, mock.Anything).Return(&[]users.Domain{usersDomain}, nil).Once()
		result, err := usersService.GetAllUser()
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		mockUserRepository.On("GetAllUser", mock.Anything, mock.Anything).Return(&[]users.Domain{usersDomain}, assert.AnError).Once()
		result, err := usersService.GetAllUser()
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockPersonalDataRepository.On("Update", mock.Anything, mock.Anything).Return(&personal_data.Domain{}, nil).Once()
		inputUser := users.Domain{
			Email:          "testuser@mail.com",
			Password:       "testUser",
			AvatarUrl:      "ava.png",
			Gender:         "male",
			PersonalDataID: 1,
			Age:            25,
			Weight:         50,
			Height:         160,
			ActivityType:   5,
		}
		inputPersonalData := personal_data.Domain{
			ID:       1,
			Calories: 1000,
			Weight:   50,
			Height:   160,
		}
		result, err := usersService.EditUser(inputUser.ID, &inputUser, &inputPersonalData)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test Users", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&usersDomain, assert.AnError).Once()
		mockPersonalDataRepository.On("Update", mock.Anything, mock.Anything).Return(&personal_data.Domain{}, nil).Once()
		inputUser := users.Domain{
			Email:          "testuser@mail.com",
			Password:       "testUser",
			AvatarUrl:      "ava.png",
			Gender:         "male",
			PersonalDataID: 1,
			Age:            25,
			Weight:         50,
			Height:         160,
			ActivityType:   5,
		}
		inputPersonalData := personal_data.Domain{
			ID:       1,
			Calories: 1000,
			Weight:   50,
			Height:   160,
		}
		result, err := usersService.EditUser(inputUser.ID, &inputUser, &inputPersonalData)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
	t.Run("Invalid Test Personal Data", func(t *testing.T) {
		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(&usersDomain, nil).Once()
		mockPersonalDataRepository.On("Update", mock.Anything, mock.Anything).Return(&personalDataDomain, assert.AnError).Once()
		inputUser := users.Domain{
			Email:          "testuser@mail.com",
			Password:       "testUser",
			AvatarUrl:      "ava.png",
			Gender:         "male",
			PersonalDataID: 1,
			Age:            25,
			Weight:         50,
			Height:         160,
			ActivityType:   5,
		}
		inputPersonalData := personal_data.Domain{
			ID:       1,
			Calories: 1000,
			Weight:   50,
			Height:   160,
		}
		_, err := usersService.EditUser(inputUser.ID, &inputUser, &inputPersonalData)
		assert.Nil(t, err)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockUserRepository.On("Delete", mock.Anything, mock.Anything).Return("Delete User Success", nil).Once()
		mockPersonalDataRepository.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
		result, err := usersService.DeleteUser(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test Personal Data", func(t *testing.T) {
		mockUserRepository.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError).Once()
		mockPersonalDataRepository.On("Delete", mock.Anything, mock.Anything).Return(assert.AnError).Once()
		result, err := usersService.DeleteUser(1)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
}
