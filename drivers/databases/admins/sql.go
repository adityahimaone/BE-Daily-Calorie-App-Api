package admins

import (
	"Daily-Calorie-App-API/business/admins"
	"gorm.io/gorm"
)

type repositoryAdmin struct {
	DB *gorm.DB
}

func NewRepositoryAdminMySQL(db *gorm.DB) admins.Repository {
	return &repositoryAdmin{
		DB: db,
	}
}

func (repository repositoryAdmin) Insert(admin *admins.Domain) (*admins.Domain, error) {
	recordAdmin := fromDomain(*admin)
	if err := repository.DB.Create(&recordAdmin).Error; err != nil {
		return &admins.Domain{}, err
	}
	result := recordAdmin.toDomain()
	return &result, nil
}

func (repository repositoryAdmin) GetAdminByUsername(username string) (*admins.Domain, error) {
	recordAdmin := Admin{}
	if err := repository.DB.Where("username = ?", username).First(&recordAdmin).Error; err != nil {
		return &admins.Domain{}, err
	}
	result := recordAdmin.toDomain()
	return &result, nil
}
