package foods

import (
	"Daily-Calorie-App-API/business/foods"
	"gorm.io/gorm"
)

type repositoryFoods struct {
	DB *gorm.DB
}

func NewRepositoryFood(db *gorm.DB) foods.Repository {
	return &repositoryFoods{
		DB: db,
	}
}

func (repository repositoryFoods) Insert(food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	if err := repository.DB.Create(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}

	if err := repository.DB.First(&recordFood).Where("id = ?", recordFood.ID).Error; err != nil {
		return &foods.Domain{}, err
	}

	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) Update(id int, food *foods.Domain) (*foods.Domain, error) {
	recordFood := fromDomain(*food)
	recordFood.ID = uint(id)
	if err := repository.DB.Where("id = ?", id).Updates(&recordFood).Error; err != nil {
		return &foods.Domain{}, err
	}

	if err := repository.DB.Joins("NutritionInfo").Find(&recordFood, "foods.id = ?", id).Error; err != nil {
		return &foods.Domain{}, err
	}

	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) Delete(id int) (string, error) {
	recordFood := Foods{}

	if err := repository.DB.Where("id = ?", id).First(&recordFood).Error; err != nil {
		return "", err
	}

	if err := repository.DB.Delete(&recordFood, "foods.id = ?", id).Error; err != nil {
		return "", err
	}

	result := "Delete Food Successfully"
	return result, nil
}

func (repository repositoryFoods) GetFoodByID(foodID int) (*foods.Domain, error) {
	recordFood := Foods{}
	if err := repository.DB.First(&recordFood, "foods.id = ?", foodID).Error; err != nil {
		return &foods.Domain{}, err
	}
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) GetFoodByName(name string) (*foods.Domain, error) {
	recordFood := Foods{}
	if err := repository.DB.Joins("NutritionInfo").Find(&recordFood, "foods.title LIKE ?", "%"+name+"%").Error; err != nil {
		return &foods.Domain{}, err
	}
	result := recordFood.toDomain()
	return &result, nil
}

func (repository repositoryFoods) GetAllFood() (*[]foods.Domain, error) {
	var recordFoods []Foods
	if err := repository.DB.Joins("NutritionInfo").Find(&recordFoods).Error; err != nil {
		return &[]foods.Domain{}, err
	}
	result := toDomainArray(recordFoods)
	return &result, nil
}
