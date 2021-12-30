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
	if err := repository.DB.Joins("NutritionInfo").First(&recordFood).Where("id = ?", recordFood.ID).Error; err != nil {
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
	if err := repository.DB.Joins("NutritionInfo").Find(&recordFood, "foods.id = ?", id).Error; err != nil {
		return nil, err
	}
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) DeleteFood(id int) (string, error) {
	recordFood := Foods{}
	if err := repository.DB.Delete(&recordFood, "foods.id = ?", id).Error; err != nil {
		return "", err
	}
	log.Println(recordFood)
	result := "Delete Food Successfully"
	return result, nil
}

func (repository repositoryFoods) GetFood(foodID int) (*foods.Domain, error) {
	recordFood := Foods{}
	if err := repository.DB.Joins("NutritionInfo").First(&recordFood, "foods.id = ?", foodID).Error; err != nil {
		return nil, err
	}
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) GetFoods() (*[]foods.Domain, error) {
	var recordFoods []Foods
	if err := repository.DB.Joins("NutritionInfo").Find(&recordFoods).Error; err != nil {
		return nil, err
	}
	result := toDomainArray(recordFoods)
	return &result, nil
}
