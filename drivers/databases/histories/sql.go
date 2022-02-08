package histories

import (
	"Daily-Calorie-App-API/business/histories"
	"github.com/kr/pretty"
	"gorm.io/gorm"
	"log"
)

type repositoryHistories struct {
	DB *gorm.DB
}

func NewRepositoryHistories(db *gorm.DB) histories.Repository {
	return &repositoryHistories{
		DB: db,
	}
}

func (repository repositoryHistories) Insert(history *histories.Domain) (*histories.Domain, error) {
	recordHistory := fromDomain(*history)
	if err := repository.DB.Create(&recordHistory).Error; err != nil {
		return &histories.Domain{}, err
	}
	if err := repository.DB.Joins("User").Joins("Food").Where("id = ?", &recordHistory.ID).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := recordHistory.toDomain()
	return &result, nil
}

func (repository repositoryHistories) Update(id int, history *histories.Domain) (*histories.Domain, error) {
	recordHistory := fromDomain(*history)
	if err := repository.DB.Where("id = ?", id).Save(&recordHistory).Error; err != nil {
		return &histories.Domain{}, err
	}
	if err := repository.DB.Joins("User").Joins("Food").Where("id = ?", &recordHistory.ID).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := recordHistory.toDomain()
	return &result, nil
}

func (repository repositoryHistories) GetHistoriesByUserIDandDate(userID int, date string) (*histories.Domain, error) {
	var recordHistories Histories
	if err := repository.DB.Where("user_id = ? AND date = ?", userID, date).Find(&recordHistories).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := recordHistories.toDomain()
	return &result, nil
}

func (repository repositoryHistories) GetAllHistoryByUserID(userID int) (*[]histories.Domain, error) {
	var recordHistories []Histories
	if err := repository.DB.Joins("User").Preload("HistoriesDetails").Preload("HistoriesDetails.Food").Order("date desc").Find(&recordHistories, "user_id = ?", userID).Error; err != nil {
		return &[]histories.Domain{}, err
	}
	log.Println(pretty.Sprint(recordHistories))
	result := toDomainArray(recordHistories)
	return &result, nil
}

func (repository repositoryHistories) GetHistoriesByID(id int) (*histories.Domain, error) {
	var recordHistories Histories
	if err := repository.DB.Where("id = ?", id).Find(&recordHistories).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := recordHistories.toDomain()
	return &result, nil
}
