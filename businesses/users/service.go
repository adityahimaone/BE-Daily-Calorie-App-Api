package users

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/businesses"
	"Daily-Calorie-App-API/businesses/personaldata"
	"Daily-Calorie-App-API/helpers"
	"log"
)

type serviceUsers struct {
	userRepository         Repository
	personaldataRepository personaldata.Repository
	jwtAuth                *auth.ConfigJWT
}

func NewService(repositoryUser Repository, repositoryPersonalData personaldata.Repository, jwtAuth *auth.ConfigJWT) Service {
	return &serviceUsers{
		userRepository:         repositoryUser,
		personaldataRepository: repositoryPersonalData,
		jwtAuth:                jwtAuth,
	}
}

func (service *serviceUsers) RegisterUser(userData *Domain, personalData *personaldata.Domain) (*Domain, *personaldata.Domain, error) {
	newPersonalData, _ := service.personaldataRepository.Insert(personalData)
	passwordHash, _ := helpers.PasswordHash(userData.Password)
	userData.Password = passwordHash
	userData.PersonalDataID = newPersonalData.ID
	result, err := service.userRepository.Insert(userData)
	if err != nil {
		return &Domain{}, &personaldata.Domain{}, businesses.ErrInsertData
	}
	return result, newPersonalData, err
}

func (service *serviceUsers) Login(email string, password string) (string, error) {
	user, err := service.userRepository.FindByEmail(email)
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

func (service *serviceUsers) FindByID(id int) (*Domain, *personaldata.Domain, error) {
	user, err := service.userRepository.FindByID(id)
	if err != nil {
		return &Domain{}, &personaldata.Domain{}, businesses.ErrUserNotFound
	}
	personalData, err := service.personaldataRepository.FindByID(user.PersonalDataID)
	if err != nil {
		return &Domain{}, &personaldata.Domain{}, businesses.ErrUserNotFound
	}
	//log.Println("personalData", personalData)
	return user, personalData, nil
}

func (service *serviceUsers) EmailAvailable(email string) (bool, error) {
	user, _ := service.userRepository.FindByEmail(email)
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (service *serviceUsers) GetAllUsers() (*[]Domain, error) {
	users, err := service.userRepository.GetAllUsers()
	if err != nil {
		return nil, businesses.ErrGetData
	}
	return users, nil
}

func (service serviceUsers) EditUser(id int, user *Domain, personalData *personaldata.Domain) (*Domain, *personaldata.Domain, error) {
	passwordHash, _ := helpers.PasswordHash(user.Password)
	user.Password = passwordHash
	userResult, err := service.userRepository.Update(id, user)
	log.Println("userResult", userResult)
	if err != nil {
		return &Domain{}, &personaldata.Domain{}, businesses.ErrUserNotFound
	}
	idPersonalData := userResult.PersonalDataID
	log.Println("idPersonalData", idPersonalData)
	personalDataResult, err := service.personaldataRepository.Update(idPersonalData, personalData)
	if err != nil {
		return &Domain{}, &personaldata.Domain{}, businesses.ErrUserNotFound
	}

	return userResult, personalDataResult, nil
}

func (service serviceUsers) DeleteUser(id int, user *Domain) (*Domain, error) {
	panic("implement me")
}
