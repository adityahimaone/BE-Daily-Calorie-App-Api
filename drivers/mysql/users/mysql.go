package users

import (
	"Daily-Calorie-App-API/businesses/users"
	"gorm.io/gorm"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) users.Repository {
	return &repositoryUsers{
		DB: db,
	}
}

func (repository repositoryUsers) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err:= repository.DB.Create(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) Update(id int, user *users.Domain) (*users.Domain, error) {
	panic("implement me")
}

func (repository repositoryUsers) FindByID(id int) (*users.Domain, error) {
	panic("implement me")
}

func (repository repositoryUsers) FindByEmail(email string) (*users.Domain, error) {
	panic("implement me")
}

func (repository repositoryUsers) GetAllUsers() ([]*users.Domain, error) {
	panic("implement me")
}

func (repository repositoryUsers) Delete(id int, user *users.Domain) (*users.Domain, error) {
	panic("implement me")
}
