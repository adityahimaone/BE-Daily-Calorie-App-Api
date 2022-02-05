package drivers

import (
	adminDomain "Daily-Calorie-App-API/business/admins"
	foodsDomain "Daily-Calorie-App-API/business/foods"
	historiesDomain "Daily-Calorie-App-API/business/histories"
	historiesdetailDomain "Daily-Calorie-App-API/business/historiesdetail"
	personaldataDomain "Daily-Calorie-App-API/business/personaldata"
	usersDomain "Daily-Calorie-App-API/business/users"

	adminDB "Daily-Calorie-App-API/drivers/databases/admins"
	foodsDB "Daily-Calorie-App-API/drivers/databases/foods"
	historiesDB "Daily-Calorie-App-API/drivers/databases/histories"
	historiesdetailDB "Daily-Calorie-App-API/drivers/databases/historiesdetail"
	personaldataDB "Daily-Calorie-App-API/drivers/databases/personaldata"
	usersDB "Daily-Calorie-App-API/drivers/databases/users"
	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) usersDomain.Repository {
	return usersDB.NewRepositoryUser(db)
}

func NewPersonalDataRepository(db *gorm.DB) personaldataDomain.Repository {
	return personaldataDB.NewRepositoryPersonalData(db)
}

func NewFoodRepository(db *gorm.DB) foodsDomain.Repository {
	return foodsDB.NewRepositoryFood(db)
}

func NewAdminRepository(db *gorm.DB) adminDomain.Repository {
	return adminDB.NewRepositoryAdminMySQL(db)
}

func NewHistoriesRepository(db *gorm.DB) historiesDomain.Repository {
	return historiesDB.NewRepositoryHistories(db)
}

func NewHistoriesDetailRepository(db *gorm.DB) historiesdetailDomain.Repository {
	return historiesdetailDB.NewRepositoryHistoriesDetail(db)
}
