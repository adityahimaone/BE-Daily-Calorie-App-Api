package postgres

import (
	"Daily-Calorie-App-API/drivers/databases/admins"
	"Daily-Calorie-App-API/drivers/databases/foods"
	"Daily-Calorie-App-API/drivers/databases/histories"
	"Daily-Calorie-App-API/drivers/databases/histories_detail"
	"Daily-Calorie-App-API/drivers/databases/meal_plans"
	"Daily-Calorie-App-API/drivers/databases/personal_data"
	"Daily-Calorie-App-API/drivers/databases/users"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type ConfigPostgresSQL struct {
	DBHost     string
	DBUsername string
	DBPassword string
	DBDatabase string
	DBPort     string
}

func (config *ConfigPostgresSQL) IntialPostgresSQL() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBDatabase,
		config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&users.Users{})
	db.AutoMigrate(&personal_data.PersonalData{})
	db.AutoMigrate(&foods.Foods{})
	db.AutoMigrate(&admins.Admin{})
	db.AutoMigrate(&histories.Histories{})
	db.AutoMigrate(&histories_detail.HistoriesDetail{})
	db.AutoMigrate(&meal_plans.MealPlans{})
}
