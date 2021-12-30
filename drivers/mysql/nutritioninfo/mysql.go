package nutritioninfo

import (
	"Daily-Calorie-App-API/businesses/nutritioninfo"
	"gorm.io/gorm"
)

type repositoryNutritionInfo struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) nutritioninfo.Repository {
	return &repositoryNutritionInfo{
		DB: db,
	}
}

func (repository repositoryNutritionInfo) Create(nutritionInfo *nutritioninfo.Domain) (*nutritioninfo.Domain, error) {
	recordNutrition := fromDomain(*nutritionInfo)
	if err := repository.DB.Create(&recordNutrition).Error; err != nil {
		return &nutritioninfo.Domain{}, err
	}
	result, _ := toDomain(recordNutrition)
	return &result, nil
}

func (repository repositoryNutritionInfo) Update(id int, nutritionInfo *nutritioninfo.Domain) (*nutritioninfo.Domain, error) {
	recordNutrition := fromDomain(*nutritionInfo)
	if err := repository.DB.Model(&recordNutrition).Where("id = ?", id).Updates(&recordNutrition).Error; err != nil {
		return &nutritioninfo.Domain{}, err
	}
	result, _ := toDomain(recordNutrition)
	return &result, nil
}

func (repository repositoryNutritionInfo) Delete(id int) error {
	recordNutrition := NutritionInfo{}
	if err := repository.DB.Delete(&recordNutrition, "nutrition_infos.id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (repository repositoryNutritionInfo) GetByID(id int) (*nutritioninfo.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (repository repositoryNutritionInfo) GetAll() ([]nutritioninfo.Domain, error) {
	//TODO implement me
	panic("implement me")
}
