package histories_detail

import (
	"Daily-Calorie-App-API/business"
)

type serviceHistoriesDetail struct {
	historiesdetailRepository Repository
}

func NewHistoriesDetailService(historiesdetailRepository Repository) Service {
	return &serviceHistoriesDetail{
		historiesdetailRepository: historiesdetailRepository,
	}
}

func (service serviceHistoriesDetail) Insert(historiesDetail *Domain) (*Domain, error) {
	result, err := service.historiesdetailRepository.Insert(historiesDetail)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service serviceHistoriesDetail) Delete(id int) (*Domain, error) {
	result, err := service.historiesdetailRepository.Delete(id)
	if err != nil {
		return &Domain{}, business.ErrNotFound
	}
	return result, nil
}

func (service serviceHistoriesDetail) GetAllHistoriesDetailByHistoriesID(historiesID int) (*[]Domain, error) {
	result, err := service.historiesdetailRepository.GetAllHistoriesDetailByHistoriesID(historiesID)
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}

func (service serviceHistoriesDetail) SumCalories(historiesID int) (float64, error) {
	result, err := service.historiesdetailRepository.SumCalories(historiesID)
	if err != nil {
		return 0, err
	}
	return result, nil
}
