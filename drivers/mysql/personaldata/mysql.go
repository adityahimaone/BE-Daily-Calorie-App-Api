package personaldata

import (
	"Daily-Calorie-App-API/businesses/personaldata"
	"gorm.io/gorm"
)

type repositoryPersonalData struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) personaldata.Repository {
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

func (repository repositoryPersonalData) Update(personaldata *personaldata.Domain) (*personaldata.Domain, error) {
	panic("implement me")
}

func (repository repositoryPersonalData) Delete(personaldata *personaldata.Domain) (*personaldata.Domain, error) {
	panic("implement me")
}

func (repository repositoryPersonalData) FindByID(id int) (*personaldata.Domain, error) {
	panic("implement me")
}
