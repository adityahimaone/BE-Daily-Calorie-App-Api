package historiesdetail

import (
	"Daily-Calorie-App-API/business/historiesdetail"
	"gorm.io/gorm"
)

type repositoryHistoriesDetail struct {
	DB *gorm.DB
}

func NewRepositoryHistoriesDetail(db *gorm.DB) historiesdetail.Repository {
	return &repositoryHistoriesDetail{
		DB: db,
	}
}

func (repository repositoryHistoriesDetail) Insert(historiesDetail *historiesdetail.Domain) (*historiesdetail.Domain, error) {
	recordHistoriesDetail := fromDomain(*historiesDetail)
	if err := repository.DB.Save(&recordHistoriesDetail).Error; err != nil {
		return nil, err
	}
	result := recordHistoriesDetail.toDomain()
	return &result, nil
}

func (repository repositoryHistoriesDetail) Update(id int, historiesDetail *historiesdetail.Domain) (*historiesdetail.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (repository repositoryHistoriesDetail) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
