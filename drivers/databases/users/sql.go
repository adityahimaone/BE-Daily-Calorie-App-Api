package users

import (
	"Daily-Calorie-App-API/business/users"
	"gorm.io/gorm"
	"log"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) users.Repository {
	return &repositoryUsers{
		DB: db,
	}
}

func (repository repositoryUsers) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Create(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}

	if err := repository.DB.Joins("PersonalData").First(&recordUser).Where("id = ?", recordUser.ID).Error; err != nil {
		return &users.Domain{}, err
	}

	result := recordUser.toDomain()
	return &result, nil
}

func (repository repositoryUsers) Update(id int, user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Where("id = ?", id).Updates(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}

	if err := repository.DB.Joins("PersonalData").Find(&recordUser).Where("users.id = ?", id).Error; err != nil {
		return &users.Domain{}, err
	}

	result := recordUser.toDomain()
	return &result, nil
}

func (repository repositoryUsers) GetUserByID(id int) (*users.Domain, error) {
	recordUser := Users{}
	if err := repository.DB.Joins("PersonalData").First(&recordUser, "users.id", id).Error; err != nil {
		return &users.Domain{}, err
	}
	result := recordUser.toDomain()
	log.Println(result, "repo")
	return &result, nil
}

func (repository repositoryUsers) GetUserByEmail(email string) (*users.Domain, error) {
	recordUser := Users{}
	if err := repository.DB.Where("email = ?", email).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := recordUser.toDomain()
	return &result, nil
}

func (repository repositoryUsers) GetAllUser() (*[]users.Domain, error) {
	var recordUsers []Users
	if err := repository.DB.Joins("PersonalData").Find(&recordUsers).Error; err != nil {
		return &[]users.Domain{}, err
	}
	result := toDomainArray(recordUsers)
	return &result, nil
}

func (repository repositoryUsers) Delete(id int) (string, error) {
	recordUser := Users{}

	if err := repository.DB.Where("id = ?", id).First(&recordUser).Error; err != nil {
		return "", err
	}

	if err := repository.DB.Where("id = ?", id).Delete(&recordUser).Error; err != nil {
		return "", err
	}
	result := "Delete User Success"
	return result, nil
}
