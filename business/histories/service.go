package histories

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	"Daily-Calorie-App-API/business/foods"
	"Daily-Calorie-App-API/business/historiesdetail"
	"Daily-Calorie-App-API/business/users"
	"log"
	"time"
)

type serviceHistories struct {
	historiesRepository              Repository
	usersService                     users.Service
	foodsService                     foods.Service
	historiesDetailRepositoryService historiesdetail.Service
	jwtAuth                          *auth.ConfigJWT
}

func NewHistoriesService(historiesRepository Repository, serviceUsers users.Service, serviceFoods foods.Service, serviceHistoriesDetail historiesdetail.Service, jwtAuth *auth.ConfigJWT) Service {
	return &serviceHistories{
		historiesRepository:              historiesRepository,
		usersService:                     serviceUsers,
		foodsService:                     serviceFoods,
		historiesDetailRepositoryService: serviceHistoriesDetail,
		jwtAuth:                          jwtAuth,
	}
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
	// log exist
	log.Println(exist, "exist")
	detailHistories := historiesdetail.Domain{}
	detailHistories.FoodID = food.ID
	if exist.ID != 0 {
		histories.TotalCalories = exist.TotalCalories + food.Calories
		histories.TotalFood = exist.TotalFood + 1
		histories.ID = exist.ID
		log.Println(histories, "histories update")
		result, err = service.historiesRepository.Update(exist.ID, histories)
		// log result
		log.Println(result, "result update")
		if err != nil {
			return &Domain{}, err
		}
		detailHistories.HistoriesID = exist.ID
		_, err = service.historiesDetailRepositoryService.Insert(&detailHistories)
		if err != nil {
			return nil, err
		}
	} else {
		histories.TotalCalories = food.Calories
		histories.TotalFood = 1
		log.Println(histories, "histories inserted")
		result, err = service.historiesRepository.Insert(histories)
		// log result
		log.Println(result, "result inserted")
		if err != nil {
			return &Domain{}, err
		}

		detailHistories.HistoriesID = result.ID
		_, err = service.historiesDetailRepositoryService.Insert(&detailHistories)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (service serviceHistories) GetAllHistoryByUserID(userID int) (*[]Domain, error) {
	//TODO implement me
	panic("implement me")
}
