package historiesdetail

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
	historiesdetail, err := service.historiesdetailRepository.Insert(historiesDetail)
	if err != nil {
		return &Domain{}, err
	}
	return historiesdetail, nil
}

func (service serviceHistoriesDetail) Update(id int, historiesDetail *Domain) (*Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (service serviceHistoriesDetail) Delete(id int) error {
	//TODO implement me
	panic("implement me")
}
