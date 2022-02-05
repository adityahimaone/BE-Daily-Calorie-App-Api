package personaldata

import (
	"Daily-Calorie-App-API/business/personaldata"
	"gorm.io/gorm"
)

type repositoryPersonalData struct {
	DB *gorm.DB
}

func NewRepositoryPersonalData(db *gorm.DB) personaldata.Repository {
	return &repositoryPersonalData{
		DB: db,
	}
}

func (repository repositoryPersonalData) Insert(personalData *personaldata.Domain) (*personaldata.Domain, error) {
	recordPersonalData := fromDomain(*personalData)
	if err := repository.DB.Create(&recordPersonalData).Error; err != nil {
		return &personaldata.Domain{}, err
	}
	result := toDomain(recordPersonalData)
	return &result, nil
}

func (repository repositoryPersonalData) Update(id int, personalData *personaldata.Domain) (*personaldata.Domain, error) {
	recordPersonalData := fromDomain(*personalData)
	if err := repository.DB.Model(&recordPersonalData).Where("id = ?", id).Updates(&recordPersonalData).Error; err != nil {
		return &personaldata.Domain{}, err
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

func (repository repositoryPersonalData) GetPersonalDataByID(id int) (*personaldata.Domain, error) {
	recordPersonalData := PersonalData{}
	if err := repository.DB.First(&recordPersonalData, id).Error; err != nil {
		return &personaldata.Domain{}, err
	}
	result := toDomain(recordPersonalData)
	return &result, nil
}
