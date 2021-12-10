package users

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/businesses/personaldata"
	"Daily-Calorie-App-API/helpers"
	"Daily-Calorie-App-API/businesses"
)

type serviceUsers struct {
	userRepository Repository
	personaldataRepository personaldata.Repository
	jwtAuth   *auth.ConfigJWT
}

func NewService(repositoryUser Repository, repositoryPersonalData personaldata.Repository, jwtAuth *auth.ConfigJWT) Service {
	return &serviceUsers{
		userRepository: repositoryUser,
		personaldataRepository: repositoryPersonalData,
		jwtAuth:        jwtAuth,
	}
}

func (service *serviceUsers) RegisterUser(userData *Domain, personalData *personaldata.Domain) (*Domain, error) {
	newPersonalData, _ := service.personaldataRepository.Insert(personalData)
	passwordHash, _ := helpers.PasswordHash(userData.Password)
	userData.Password = passwordHash
	userData.PersonalDataID = newPersonalData.ID
	valid, _ := service.EmailAvailable(userData.Email)
	if valid == true {
		result, err := service.userRepository.Insert(userData)
		if err != nil {
			return &Domain{}, businesses.ErrInsertData
		}
		return result, err
	}
	return &Domain{}, businesses.ErrDuplicateData
}

func (service *serviceUsers) Login(email string, password string) (string, error) {
	user, err := service.userRepository.FindByEmail(email )
	if err != nil {
		return "User Not Found", businesses.ErrUserNotFound
	}
	if user.ID == 0 {
		return "User Not Found", businesses.ErrUserNotFound
	}
	if !helpers.ValidateHash(password, user.Password) {
		return "Error Validate Hash", businesses.ErrInvalidLogin
	}
	token := service.jwtAuth.GenerateToken(user.ID)
	return token, err
}

func (service *serviceUsers) FindByID(id int) (*Domain, error) {
	user, err := service.FindByID(id)
	if err != nil {
		return &Domain{}, businesses.ErrUserNotFound
	}
	return user, nil
}

func (service *serviceUsers) EmailAvailable(email string) (bool, error) {
	user, _ := service.userRepository.FindByEmail(email)
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (service *serviceUsers) GetAllUsers() ([]*Domain, error) {
	panic("implement me")
}

func (service serviceUsers) EditUser(id int, user *Domain) (*Domain, error) {
	panic("implement me")
}

func (service serviceUsers) DeleteUser(id int, user *Domain) (*Domain, error) {
	panic("implement me")
}


