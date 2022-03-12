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

func (r repositoryAdmin) Insert(admin *admins.Domain) (*admins.Domain, error) {
	recordAdmin := fromDomain(*admin)
	if err := r.DB.Create(&recordAdmin).Error; err != nil {
		return &admins.Domain{}, err
	}
	result := recordAdmin.toDomain()
	return &result, nil
}

func (r repositoryAdmin) GetAdminByUsername(username string) (*admins.Domain, error) {
	recordAdmin := Admin{}
	if err := r.DB.Where("username = ?", username).First(&recordAdmin).Error; err != nil {
		return &admins.Domain{}, err
	}
	result := recordAdmin.toDomain()
	return &result, nil
}
