package users

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business"
	"Daily-Calorie-App-API/business/personal_data"
	"Daily-Calorie-App-API/helpers"
	"math"
)

type serviceUsers struct {
	userRepository         Repository
	personaldataRepository personal_data.Repository
	jwtAuth                *auth.ConfigJWT
}

func NewUserService(repositoryUser Repository, repositoryPersonalData personal_data.Repository, jwtAuth *auth.ConfigJWT) Service {
	return &serviceUsers{
		userRepository:         repositoryUser,
		personaldataRepository: repositoryPersonalData,
		jwtAuth:                jwtAuth,
	}
}

func (service *serviceUsers) RegisterUser(userData *Domain, personalData *personal_data.Domain) (*Domain, error) {
	countCalories, err := service.CountCalories(userData, personalData)
	if err != nil {
		return nil, err
	}
	personalData.Calories = countCalories
	newPersonalData, err := service.personaldataRepository.Insert(personalData)
	if err != nil {
		return &Domain{}, business.ErrInsertData
	}
	passwordHash, _ := helpers.PasswordHash(userData.Password)
	userData.Password = passwordHash
	userData.PersonalDataID = newPersonalData.ID
	result, err := service.userRepository.Insert(userData)
	if err != nil {
		return &Domain{}, business.ErrInsertData
	}
	return result, err
}

func (service *serviceUsers) CountCalories(userData *Domain, personalData *personal_data.Domain) (float64, error) {
	activityTypeValue := userData.ActivityType
	weight := float64(personalData.Weight)
	height := float64(personalData.Height)
	age := float64(userData.Age)
	gender := userData.Gender
	calories := 0.0

	if gender == "male" {
		calories = (10 * weight) + (6.25 * height) - (5*age)*activityTypeValue
	} else {
		calories = ((10 * weight) + (6.25 * height) - (5 * age) - 161) * activityTypeValue
	}

	result := math.Round(calories*100) / 100
	return result, nil
}

func (service *serviceUsers) Login(email string, password string) (string, error) {
	user, err := service.userRepository.GetUserByEmail(email)
	if err != nil {
		return "User Not Found", business.ErrUserNotFound
	}
	if user.ID == 0 {
		return "User Not Found", business.ErrUserNotFound
	}
	if !helpers.ValidateHash(password, user.Password) {
		return "Error Validate Hash", business.ErrInvalidLogin
	}
	token := service.jwtAuth.GenerateToken(user.ID, "user")
	return token, err
}

func (service *serviceUsers) GetUserByID(id int) (*Domain, error) {
	user, err := service.userRepository.GetUserByID(id)
	if err != nil {
		return &Domain{}, business.ErrUserNotFound
	}
	return user, nil
}

func (service *serviceUsers) EmailAvailable(email string) (bool, error) {
	user, _ := service.userRepository.GetUserByEmail(email)
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (service *serviceUsers) GetAllUser() (*[]Domain, error) {
	users, err := service.userRepository.GetAllUser()
	if err != nil {
		return &[]Domain{}, business.ErrGetData
	}
	return users, nil
}

func (service serviceUsers) EditUser(id int, user *Domain, personalData *personal_data.Domain) (*Domain, error) {
	passwordHash, _ := helpers.PasswordHash(user.Password)
	user.Password = passwordHash
	userResult, err := service.userRepository.Update(id, user)
	if err != nil {
		return &Domain{}, business.ErrUserNotFound
	}
	idPersonalData := userResult.PersonalDataID
	_, err = service.personaldataRepository.Update(idPersonalData, personalData)
	if err != nil {
		return &Domain{}, business.ErrUserNotFound
	}

	return userResult, nil
}

func (service serviceUsers) DeleteUser(id int) (string, error) {
	err := service.personaldataRepository.Delete(id)
	if err != nil {
		return "Error Delete Personal Data", business.ErrDeleteData
	}
	result, err := service.userRepository.Delete(id)
	if err != nil {
		return "User Not Found", business.ErrUserNotFound
	}
	return result, nil
}
