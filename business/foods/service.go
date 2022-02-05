package foods

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business"
)

type serviceFoods struct {
	foodRepository Repository
	jwtAuth        *auth.ConfigJWT
}

func NewFoodService(repositoryFood Repository, jwtAuth *auth.ConfigJWT) Service {
	return &serviceFoods{
		foodRepository: repositoryFood,
		jwtAuth:        jwtAuth,
	}
}

func (service serviceFoods) AddFood(food *Domain) (*Domain, error) {
	result, err := service.foodRepository.Insert(food)
	if err != nil {
		return &Domain{}, business.ErrInsertData
	}
	return result, nil
}

func (service serviceFoods) EditFood(id int, food *Domain) (*Domain, error) {
	result, err := service.foodRepository.Update(id, food)
	if err != nil {
		return &Domain{}, business.ErrUpdateData
	}
	return result, nil
}

func (service serviceFoods) DeleteFood(id int) (string, error) {
	result, err := service.foodRepository.Delete(id)
	if err != nil {
		return "", business.ErrNotFound
	}
	return result, nil
}

func (service serviceFoods) GetFoodByID(id int) (*Domain, error) {
	result, err := service.foodRepository.GetFoodByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service serviceFoods) GetFoodByName(name string) (*Domain, error) {
	result, err := service.foodRepository.GetFoodByName(name)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service serviceFoods) GetAllFood() (*[]Domain, error) {
	result, err := service.foodRepository.GetAllFood()
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}
