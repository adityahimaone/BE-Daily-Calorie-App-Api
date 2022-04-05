package histories

import (
	"Daily-Calorie-App-API/business/histories"
	"gorm.io/gorm"
)

type repositoryHistories struct {
	DB *gorm.DB
}

func NewRepositoryHistories(db *gorm.DB) histories.Repository {
	return &repositoryHistories{
		DB: db,
	}
}

func (repository repositoryHistories) UpdateTotalHistoriesDetail(id int) (*histories.Domain, error) {
	var history histories.Domain
	err := repository.DB.Where("id = ?", id).First(&history).Error
	if err != nil {
		return nil, err
	}
	err = repository.DB.Model(&history).Updates(histories.Domain{
		TotalCalories: history.TotalCalories,
	}).Error
	if err != nil {
		return nil, err
	}
	return &history, nil
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

func (repository repositoryHistories) UpdateWater(id int, history *histories.Domain) (*histories.Domain, error) {
	recordHistory := fromDomain(*history)
	if err := repository.DB.Model(&recordHistory).Where("id = ?", id).UpdateColumn("water", &recordHistory.Water).Error; err != nil {
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

func (repository repositoryHistories) GetLastHistoryByUserID(userID int) (*histories.Domain, error) {
	var recordHistories Histories
	//dateTime := time.Now().Format("02-01-2006")
	if err := repository.DB.Joins("User").Preload("HistoriesDetails").Preload("HistoriesDetails.Food").Order("date asc").Last(&recordHistories, "user_id = ?", userID).Error; err != nil {
		return &histories.Domain{}, err
	}
	//log.Println(pretty.Sprint(recordHistories))
	result := recordHistories.toDomain()
	return &result, nil
}

func (repository repositoryHistories) GetAllHistoryByUserID(userID int) (*[]histories.Domain, error) {
	var recordHistories []Histories
	if err := repository.DB.Joins("User").Preload("HistoriesDetails").Preload("HistoriesDetails.Food").Order("date asc").Find(&recordHistories, "user_id = ?", userID).Error; err != nil {
		return &[]histories.Domain{}, err
	}
	result := toDomainArray(recordHistories)
	return &result, nil
}

func (repository repositoryHistories) GetHistoriesByID(id int) (*histories.Domain, error) {
	var recordHistories Histories
	if err := repository.DB.Joins("User").Preload("HistoriesDetails").Preload("HistoriesDetails.Food").Find(&recordHistories, "histories.id = ?", id).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := recordHistories.toDomain()
	return &result, nil
}

func (repository repositoryHistories) UpdateTotalCalories(id int, totalCalories float64) (*histories.Domain, error) {
	var recordHistories Histories
	if err := repository.DB.Model(&recordHistories).Where("id = ?", id).UpdateColumn("total_calories", totalCalories).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := recordHistories.toDomain()
	return &result, nil
}

func (repository repositoryHistories) UpdateTotalFood(id int, totalFood int) (*histories.Domain, error) {
	var recordHistories Histories
	if err := repository.DB.Model(&recordHistories).Where("id = ?", id).UpdateColumn("total_food", totalFood).Error; err != nil {
		return &histories.Domain{}, err
	}
	result := recordHistories.toDomain()
	return &result, nil
}
