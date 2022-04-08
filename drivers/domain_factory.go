package drivers

import (
	adminDomain "Daily-Calorie-App-API/business/admins"
	foodsDomain "Daily-Calorie-App-API/business/foods"
	foodsapiDomain "Daily-Calorie-App-API/business/foodsAPI"
	historiesDomain "Daily-Calorie-App-API/business/histories"
	historiesdetailDomain "Daily-Calorie-App-API/business/histories_detail"
	mealplansDomain "Daily-Calorie-App-API/business/meal_plans"
	personaldataDomain "Daily-Calorie-App-API/business/personal_data"
	usersDomain "Daily-Calorie-App-API/business/users"

	adminDB "Daily-Calorie-App-API/drivers/databases/admins"
	foodsDB "Daily-Calorie-App-API/drivers/databases/foods"
	historiesDB "Daily-Calorie-App-API/drivers/databases/histories"
	historiesdetailDB "Daily-Calorie-App-API/drivers/databases/histories_detail"
	mealplansDB "Daily-Calorie-App-API/drivers/databases/meal_plans"
	personaldataDB "Daily-Calorie-App-API/drivers/databases/personal_data"
	usersDB "Daily-Calorie-App-API/drivers/databases/users"
	foodAPI "Daily-Calorie-App-API/drivers/thirdparties/edamamAPI"
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
	return adminDB.NewRepositoryAdmin(db)
}

func NewHistoriesRepository(db *gorm.DB) historiesDomain.Repository {
	return historiesDB.NewRepositoryHistories(db)
}

func NewHistoriesDetailRepository(db *gorm.DB) historiesdetailDomain.Repository {
	return historiesdetailDB.NewRepositoryHistoriesDetail(db)
}

func NewFoodAPIRepository() foodsapiDomain.Repository {
	return foodAPI.NewFoodAPI()
}

func NewMealPlansRepository(db *gorm.DB) mealplansDomain.Repository {
	return mealplansDB.NewRepositoryMealPlans(db)
}
