package response

import (
	"Daily-Calorie-App-API/business/histories"
	"Daily-Calorie-App-API/drivers/databases/histories_detail"
	"time"
)

type Histories struct {
	ID               int                                `json:"id"`
	UserID           int                                `json:"user_id"`
	Fullname         string                             `json:"fullname"`
	Water            int                                `json:"water"`
	TotalCalories    float64                            `json:"total_calories"`
	TotalFood        int                                `json:"total_food"`
	Date             string                             `json:"date"`
	CreatedAt        time.Time                          `json:"created_at"`
	UpdatedAt        time.Time                          `json:"updated_at"`
	HistoriesDetails []histories_detail.HistoriesDetail `json:"histories_details"`
}

func FromDomain(domain histories.Domain) Histories {
	return Histories{
		ID:               domain.ID,
		UserID:           domain.UserID,
		Fullname:         domain.UserName,
		Water:            domain.Water,
		TotalCalories:    domain.TotalCalories,
		TotalFood:        domain.TotalFood,
		Date:             domain.Date,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		HistoriesDetails: domain.HistoriesDetail,
	}
}

func FromDomainArray(domains []histories.Domain) []Histories {
	var historiesArray []Histories
	for _, domain := range domains {
		historiesArray = append(historiesArray, FromDomain(domain))
	}
	return historiesArray
}
