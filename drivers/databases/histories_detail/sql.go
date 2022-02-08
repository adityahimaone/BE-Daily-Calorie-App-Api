package histories_detail

import (
	"Daily-Calorie-App-API/business/histories_detail"
	"gorm.io/gorm"
)

type repositoryHistoriesDetail struct {
	DB *gorm.DB
}

func NewRepositoryHistoriesDetail(db *gorm.DB) histories_detail.Repository {
	return &repositoryHistoriesDetail{
		DB: db,
	}
}

func (repository repositoryHistoriesDetail) Insert(historiesDetail *histories_detail.Domain) (*histories_detail.Domain, error) {
	recordHistoriesDetail := fromDomain(*historiesDetail)
	if err := repository.DB.Save(&recordHistoriesDetail).Error; err != nil {
		return nil, err
	}
	result := recordHistoriesDetail.toDomain()
	return &result, nil
}

func (repository repositoryHistoriesDetail) Delete(id int) error {
	recordHistoriesDetail := HistoriesDetail{}
	if err := repository.DB.Delete(&recordHistoriesDetail, id).Error; err != nil {
		return err
	}
	return nil
}

func (repository repositoryHistoriesDetail) GetAllHistoriesDetailByHistoriesID(historiesID int) (*[]histories_detail.Domain, error) {
	var recordHistoriesDetail []HistoriesDetail
	if err := repository.DB.Joins("Food").Find(&recordHistoriesDetail, "histories_id = ?", historiesID).Error; err != nil {
		return nil, err
	}
	result := toDomainArray(recordHistoriesDetail)
	return &result, nil
}
