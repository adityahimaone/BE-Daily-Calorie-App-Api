package personal_data

import (
	"Daily-Calorie-App-API/business/personal_data"
	"gorm.io/gorm"
)

type repositoryPersonalData struct {
	DB *gorm.DB
}

func NewRepositoryPersonalData(db *gorm.DB) personal_data.Repository {
	return &repositoryPersonalData{
		DB: db,
	}
}

func (repository repositoryPersonalData) Insert(personalData *personal_data.Domain) (*personal_data.Domain, error) {
	recordPersonalData := fromDomain(*personalData)
	if err := repository.DB.Create(&recordPersonalData).Error; err != nil {
		return &personal_data.Domain{}, err
	}
	result := toDomain(recordPersonalData)
	return &result, nil
}

func (repository repositoryPersonalData) Update(id int, personalData *personal_data.Domain) (*personal_data.Domain, error) {
	recordPersonalData := fromDomain(*personalData)
	if err := repository.DB.Model(&recordPersonalData).Where("id = ?", id).Updates(&recordPersonalData).Error; err != nil {
		return &personal_data.Domain{}, err
	}
	result := toDomain(recordPersonalData)
	return &result, nil
}

func (repository repositoryPersonalData) Delete(id int) error {
	recordPersonalData := PersonalData{}
	if err := repository.DB.Delete(&recordPersonalData, "personal_data.id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (repository repositoryPersonalData) GetPersonalDataByID(id int) (*personal_data.Domain, error) {
	recordPersonalData := PersonalData{}
	if err := repository.DB.First(&recordPersonalData, id).Error; err != nil {
		return &personal_data.Domain{}, err
	}
	result := toDomain(recordPersonalData)
	return &result, nil
}
