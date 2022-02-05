package historiesdetail

import (
	"Daily-Calorie-App-API/business/historiesdetail"
	"Daily-Calorie-App-API/drivers/databases/foods"
	"gorm.io/gorm"
	"time"
)

type HistoriesDetail struct {
	gorm.Model
	ID          uint        `gorm:"primary_key"`
	HistoriesID uint        `gorm:"not null"`
	FoodID      uint        `gorm:"not null"`
	Food        foods.Foods `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func fromDomain(domain historiesdetail.Domain) HistoriesDetail {
	return HistoriesDetail{
		ID:          uint(domain.ID),
		HistoriesID: uint(domain.HistoriesID),
		FoodID:      uint(domain.FoodID),
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func (record *HistoriesDetail) toDomain() historiesdetail.Domain {
	return historiesdetail.Domain{
		ID:          int(record.ID),
		HistoriesID: int(record.HistoriesID),
		FoodID:      int(record.FoodID),
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}
