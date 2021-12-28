package foods

import (
	"Daily-Calorie-App-API/businesses/foods"
	"gorm.io/gorm"
	"log"
)

type repositoryFoods struct {
	DB *gorm.DB
}

func NewRepositoryFoodMySQL(db *gorm.DB) foods.Repository {
	return &repositoryFoods{
		DB: db,
	}
}

func (repository repositoryFoods) InsertFood(food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Create(&recordFood).Error; err != nil {
		return nil, err
	}
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) UpdateFood(id int, food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Model(&recordFood).Where("id = ?", id).Updates(&recordFood).Error; err != nil {
		return nil, err
	}
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) DeleteFood(id int) (*foods.Domain, error) {
	recordFood := Foods{}
	if err := repository.DB.Delete(&recordFood).Where("id = ?", id).Error; err != nil {
		return nil, err
	}
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) GetFood(id int) (*foods.Domain, error) {
	recordFood := Foods{}
	if err := repository.DB.Joins("NutritionInfo").First(&recordFood).Where("id = ?", id).Error; err != nil {
		return nil, err
	}
	log.Println(recordFood, "recordFood")
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) GetFoods() (*[]foods.Domain, error) {
	var recordFoods []Foods
	if err := repository.DB.Find(&recordFoods).Error; err != nil {
		return nil, err
	}
	result := toDomainArray(recordFoods)
	return &result, nil
}
