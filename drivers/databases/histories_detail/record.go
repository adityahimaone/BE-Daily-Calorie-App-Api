package histories_detail

import (
	"Daily-Calorie-App-API/business/histories_detail"
	"Daily-Calorie-App-API/drivers/databases/foods"
	"gorm.io/gorm"
	"time"
)

type HistoriesDetail struct {
	gorm.Model
	ID          uint        `gorm:"primary_key" json:"ID,omitempty"`
	HistoriesID uint        `gorm:"not null" json:"historiesID,omitempty"`
	FoodID      uint        `gorm:"not null" json:"foodID,omitempty"`
	Food        foods.Foods `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"food"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func fromDomain(domain histories_detail.Domain) HistoriesDetail {
	return HistoriesDetail{
		ID:          uint(domain.ID),
		HistoriesID: uint(domain.HistoriesID),
		FoodID:      uint(domain.FoodID),
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (record *HistoriesDetail) toDomain() histories_detail.Domain {
	return histories_detail.Domain{
		ID:           int(record.ID),
		HistoriesID:  int(record.HistoriesID),
		FoodID:       int(record.FoodID),
		FoodTitle:    record.Food.Title,
		FoodImage:    record.Food.ImgURL,
		FoodCalories: record.Food.Calories,
		FoodCrabs:    record.Food.Carbs,
		FoodProtein:  record.Food.Protein,
		FoodFat:      record.Food.Fat,
		CreatedAt:    record.CreatedAt,
		UpdatedAt:    record.UpdatedAt,
	}
}

func toDomainArray(record []HistoriesDetail) []histories_detail.Domain {
	var domains []histories_detail.Domain
	for _, v := range record {
		domains = append(domains, v.toDomain())
	}
	return domains
}
