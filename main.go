package main

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	_middlewareLog "Daily-Calorie-App-API/app/middleware/log"
	_routes "Daily-Calorie-App-API/app/routes"

	_serviceAdmins "Daily-Calorie-App-API/business/admins"
	_serviceFood "Daily-Calorie-App-API/business/foods"
	_serviceFoodAPI "Daily-Calorie-App-API/business/foodsAPI"
	_serviceHistories "Daily-Calorie-App-API/business/histories"
	_serviceHistoriesDetail "Daily-Calorie-App-API/business/histories_detail"
	_serviceMealplans "Daily-Calorie-App-API/business/meal_plans"
	_serviceUsers "Daily-Calorie-App-API/business/users"

	_controllerAdmin "Daily-Calorie-App-API/controllers/admins"
	_controllerFood "Daily-Calorie-App-API/controllers/foods"
	_controllerFoodsAPI "Daily-Calorie-App-API/controllers/foodsAPI"
	_coontrollerHistories "Daily-Calorie-App-API/controllers/histories"
	_controllerHistoriesDetail "Daily-Calorie-App-API/controllers/histories_detail"
	_mealplanController "Daily-Calorie-App-API/controllers/meal_plans"
	_controllerUser "Daily-Calorie-App-API/controllers/users"

	_driverFactory "Daily-Calorie-App-API/drivers"
	_dbPostgres "Daily-Calorie-App-API/drivers/postgres"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("app-config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func main() {
	//configDB := _dbDriver.ConfigDB{
	//	DBUsername: viper.GetString(`database.user`),
	//	DBPassword: viper.GetString(`database.pass`),
	//	DBHost:     viper.GetString(`database.host`),
	//	DBPort:     viper.GetString(`database.port`),
	//	DBDatabase: viper.GetString(`database.name`),
	//}

	configPostgres := _dbPostgres.ConfigPostgresSQL{
		DBHost:     viper.GetString(`postgres.host`),
		DBUsername: viper.GetString(`postgres.user`),
		DBPassword: viper.GetString(`postgres.pass`),
		DBDatabase: viper.GetString(`postgres.name`),
		DBPort:     viper.GetString(`postgres.port`),
	}

	configJWT := auth.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	//db := configDB.IntialDB()
	dbPostgres := configPostgres.IntialPostgresSQL()
	//_dbDriver.MigrateDB(db)
	_dbPostgres.MigrateDB(dbPostgres)
	e := echo.New()

	//factory of domain
	personaldataRepository := _driverFactory.NewPersonalDataRepository(dbPostgres)

	userRepository := _driverFactory.NewUserRepository(dbPostgres)
	userService := _serviceUsers.NewUserService(userRepository, personaldataRepository, &configJWT)
	userController := _controllerUser.NewController(userService)

	foodRepository := _driverFactory.NewFoodRepository(dbPostgres)
	foodService := _serviceFood.NewFoodService(foodRepository, &configJWT)
	foodController := _controllerFood.NewController(foodService)

	adminRepository := _driverFactory.NewAdminRepository(dbPostgres)
	adminService := _serviceAdmins.NewServiceAdmin(adminRepository, &configJWT)
	adminController := _controllerAdmin.NewController(adminService)

	historiesdetailRepository := _driverFactory.NewHistoriesDetailRepository(dbPostgres)
	historiesdetailService := _serviceHistoriesDetail.NewHistoriesDetailService(historiesdetailRepository, foodService)
	historiesdetailController := _controllerHistoriesDetail.NewController(historiesdetailService)

	historiesRepository := _driverFactory.NewHistoriesRepository(dbPostgres)
	historiesService := _serviceHistories.NewHistoriesService(historiesRepository, userService, foodService, historiesdetailService, &configJWT)
	historiesController := _coontrollerHistories.NewController(historiesService)

	foodapiRepository := _driverFactory.NewFoodAPIRepository()
	foodapiService := _serviceFoodAPI.NewFoodAPIService(foodapiRepository, &configJWT)
	foodapiController := _controllerFoodsAPI.NewController(foodapiService)

	mealplansRepository := _driverFactory.NewMealPlansRepository(dbPostgres)
	mealplansService := _serviceMealplans.NewMealPlansService(mealplansRepository, foodapiService, &configJWT)
	mealplansController := _mealplanController.NewController(mealplansService)

	// initial of route
	routesInit := _routes.HandlerList{
		JWTMiddleware:             configJWT.Init(),
		UserController:            *userController,
		FoodController:            *foodController,
		AdminController:           *adminController,
		HistoriesController:       *historiesController,
		HistoriesDetailController: *historiesdetailController,
		FoodAPIController:         *foodapiController,
		MealPlansController:       *mealplansController,
	}

	routesInit.RouteRegister(e)
	_middlewareLog.LogMiddleware(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
