package main

import (
	"Daily-Calorie-App-API/app/middleware/auth"
	_middlewareLog "Daily-Calorie-App-API/app/middleware/log"
	"Daily-Calorie-App-API/app/routes"
	foods2 "Daily-Calorie-App-API/businesses/foods"
	_serviceUsers "Daily-Calorie-App-API/businesses/users"
	foods3 "Daily-Calorie-App-API/controllers/foods"
	_controllerUser "Daily-Calorie-App-API/controllers/users"
	mysqlDriver "Daily-Calorie-App-API/drivers/mysql"
	"Daily-Calorie-App-API/drivers/mysql/foods"
	"Daily-Calorie-App-API/drivers/mysql/nutritioninfo"
	_repositoryPersonalData "Daily-Calorie-App-API/drivers/mysql/personaldata"
	_repositoryUsers "Daily-Calorie-App-API/drivers/mysql/users"
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
	configDB := mysqlDriver.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}

	configJWT := auth.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	db := configDB.IntialDB()
	mysqlDriver.MigrateDB(db)

	e := echo.New()

	//factory of domain
	personaldataRepository := _repositoryPersonalData.NewRepositoryMySQL(db)

	userRepository := _repositoryUsers.NewRepositoryMySQL(db)
	userService := _serviceUsers.NewService(userRepository, personaldataRepository, &configJWT)
	userController := _controllerUser.NewController(userService)

	nutritioninfoRepsoitory := nutritioninfo.NewRepositoryMySQL(db)

	foodRepository := foods.NewRepositoryFoodMySQL(db)
	foodService := foods2.NewService(foodRepository, nutritioninfoRepsoitory)
	foodController := foods3.NewController(foodService)

	// initial of route
	routesInit := routes.HandlerList{
		JWTMiddleware:  configJWT.Init(),
		UserController: *userController,
		FoodController: *foodController,
	}
	routesInit.RouteRegister(e)
	_middlewareLog.LogMiddleware(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
