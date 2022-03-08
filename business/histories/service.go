package histories

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/foods"
	"Daily-Calorie-App-API/business/histories_detail"
	"Daily-Calorie-App-API/business/users"
	"log"
	"time"
)

type serviceHistories struct {
	historiesRepository    Repository
	usersService           users.Service
	foodsService           foods.Service
	historiesDetailService histories_detail.Service
	jwtAuth                *auth.ConfigJWT
}

func NewHistoriesService(historiesRepository Repository, serviceUsers users.Service, serviceFoods foods.Service, serviceHistoriesDetail histories_detail.Service, jwtAuth *auth.ConfigJWT) Service {
	return &serviceHistories{
		historiesRepository:    historiesRepository,
		usersService:           serviceUsers,
		foodsService:           serviceFoods,
		historiesDetailService: serviceHistoriesDetail,
		jwtAuth:                jwtAuth,
	}
}

func (service serviceHistories) CreateHistoriesFromAPI(histories *Domain, food *foods.Domain) (*Domain, error) {
	FindFood, err := service.foodsService.GetFoodByName(food.Title)
	if err != nil {
		return &Domain{}, err
	}
	if FindFood.ID == 0 {
		insertFood, err := service.foodsService.AddFood(food)
		if err != nil {
			return &Domain{}, err
		}
		histories.FoodID = insertFood.ID
		result, err := service.CreateHistories(histories)
		if err != nil {
			return &Domain{}, err
		}
		return result, nil
	}
	histories.FoodID = FindFood.ID
	result, err := service.CreateHistories(histories)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (service serviceHistories) CreateWater(histories *Domain) (*Domain, error) {
	user, err := service.usersService.GetUserByID(histories.UserID)
	if err != nil {
		return &Domain{}, err
	}
	result := histories
	histories.UserID = user.ID
	dateTime := time.Now().Format("02-01-2006")
	histories.Date = dateTime
	exist, err := service.historiesRepository.GetHistoriesByUserIDandDate(histories.UserID, histories.Date)
	if exist.ID != 0 {
		histories.ID = exist.ID
		result, err = service.historiesRepository.UpdateWater(exist.ID, histories)
		if err != nil {
			return &Domain{}, err
		}
	} else {
		result, err = service.historiesRepository.Insert(histories)
		if err != nil {
			return &Domain{}, err
		}
	}
	return result, nil
}

func (service serviceHistories) CreateHistories(histories *Domain) (*Domain, error) {
	user, err := service.usersService.GetUserByID(histories.UserID)
	if err != nil {
		return &Domain{}, err
	}

	food, err := service.foodsService.GetFoodByID(histories.FoodID)
	if err != nil {
		return &Domain{}, err
	}

	result := histories
	histories.UserID = user.ID
	dateTime := time.Now().Format("02-01-2006")
	histories.Date = dateTime
	exist, err := service.historiesRepository.GetHistoriesByUserIDandDate(histories.UserID, histories.Date)
	detailHistories := histories_detail.Domain{}
	detailHistories.FoodID = food.ID
	if exist.ID != 0 {
		histories.TotalCalories = exist.TotalCalories + food.Calories
		histories.TotalFood = exist.TotalFood + 1
		histories.ID = exist.ID
		result, err = service.historiesRepository.Update(exist.ID, histories)
		if err != nil {
			return &Domain{}, err
		}
		detailHistories.HistoriesID = exist.ID
		_, err = service.historiesDetailService.Insert(&detailHistories)
		if err != nil {
			return nil, err
		}
	} else {
		histories.TotalCalories = food.Calories
		histories.TotalFood = 1
		result, err = service.historiesRepository.Insert(histories)
		if err != nil {
			return &Domain{}, err
		}
		detailHistories.HistoriesID = result.ID
		_, err = service.historiesDetailService.Insert(&detailHistories)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (service serviceHistories) GetLastHistoryByUserID(userID int) (*Domain, error) {
	result, err := service.historiesRepository.GetLastHistoryByUserID(userID)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service serviceHistories) GetAllHistoryByUserID(userID int) (*[]Domain, error) {
	result, err := service.historiesRepository.GetAllHistoryByUserID(userID)
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}

func (service serviceHistories) GetHistoriesByUserIDandDate(userID int) (*Domain, error) {
	dateTime := time.Now().Format("02-01-2006")
	result, err := service.historiesRepository.GetHistoriesByUserIDandDate(userID, dateTime)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service serviceHistories) GetHistoriesByID(id int) (*Domain, error) {
	result, err := service.historiesRepository.GetHistoriesByID(id)
	if err != nil {
		return &Domain{}, err
	}
	// count histories_details
	CountHistoriesDetail := len(result.HistoriesDetail)
	log.Println(CountHistoriesDetail)
	return result, nil
}

func (service serviceHistories) DeleteHistoriesDetail(historiesDetailID int) (string, error) {
	resultDelete, err := service.historiesDetailService.Delete(historiesDetailID)
	if err != nil {
		return "Error Delete", err
	}
	resultSumCalorie, err := service.historiesDetailService.SumCalories(resultDelete.HistoriesID)
	if err != nil {
		return "Error Sum", err
	}
	_, err = service.historiesRepository.UpdateTotalCalories(resultDelete.HistoriesID, resultSumCalorie)
	if err != nil {
		return "Error Update", err
	}
	resultGetHistory, err := service.historiesRepository.GetHistoriesByID(resultDelete.HistoriesID)
	if err != nil {
		return "Error Get History", err
	}
	// count histories_details
	CountHistoriesDetail := len(resultGetHistory.HistoriesDetail)
	_, err = service.historiesRepository.UpdateTotalFood(resultDelete.HistoriesID, CountHistoriesDetail)
	if err != nil {
		return "Error Update Total Food", err
	}
	return "Success", nil
}

func (service serviceHistories) UpdateTotalCalories(historiesID int, totalCalories float64) (*Domain, error) {
	result, err := service.historiesRepository.UpdateTotalCalories(historiesID, totalCalories)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
