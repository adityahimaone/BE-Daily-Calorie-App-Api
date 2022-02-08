package histories_detail

import "Daily-Calorie-App-API/business/foods"

type serviceHistoriesDetail struct {
	historiesdetailRepository Repository
	foodsService              foods.Service
}

func NewHistoriesDetailService(historiesdetailRepository Repository, foodsService foods.Service) Service {
	return &serviceHistoriesDetail{
		historiesdetailRepository: historiesdetailRepository,
		foodsService:              foodsService,
	}
}

func (service serviceHistoriesDetail) Insert(historiesDetail *Domain) (*Domain, error) {
	result, err := service.historiesdetailRepository.Insert(historiesDetail)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service serviceHistoriesDetail) Delete(id int) (string, error) {
	err := service.historiesdetailRepository.Delete(id)
	if err != nil {
		return "Error Delete", err
	}
	return "Success Delete", nil
}

func (service serviceHistoriesDetail) GetAllHistoriesDetailByHistoriesID(historiesID int) (*[]Domain, error) {
	result, err := service.historiesdetailRepository.GetAllHistoriesDetailByHistoriesID(historiesID)
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}
