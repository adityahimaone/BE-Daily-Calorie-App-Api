package request

import "Daily-Calorie-App-API/business/histories_detail"

type HistoriesDetail struct {
	ID          int `json:"id"`
	HistoriesID int `json:"histories_id"`
}

func (request *HistoriesDetail) ToDomain() *histories_detail.Domain {
	return &histories_detail.Domain{
		ID:          request.ID,
		HistoriesID: request.HistoriesID,
	}
}
