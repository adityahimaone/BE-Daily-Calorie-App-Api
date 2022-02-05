package histories

import (
	"Daily-Calorie-App-API/business/histories"
	"Daily-Calorie-App-API/drivers/databases/historiesdetail"
	"Daily-Calorie-App-API/drivers/databases/users"
	"gorm.io/gorm"
	"time"
)

type Histories struct {
	gorm.Model
	ID               uint                              `gorm:"primary_key"`
	UserID           uint                              `gorm:"not null"`
	User             users.Users                       `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	HistoriesDetails []historiesdetail.HistoriesDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Water            int                               `gorm:"type:int"`
	TotalCalories    float64                           `gorm:"type:decimal(10,2)"`
	TotalFood        int                               `gorm:"type:int"`
	Date             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (record *Histories) toDomain() histories.Domain {
	return histories.Domain{
		ID:            int(record.ID),
		UserID:        int(record.UserID),
		Date:          record.Date,
		Water:         record.Water,
		TotalCalories: record.TotalCalories,
		TotalFood:     record.TotalFood,
		CreatedAt:     record.CreatedAt,
		UpdatedAt:     record.UpdatedAt,
	}
}

func fromDomain(domain histories.Domain) Histories {
	return Histories{
		ID:            uint(domain.ID),
		UserID:        uint(domain.UserID),
		Date:          domain.Date,
		Water:         domain.Water,
		TotalCalories: domain.TotalCalories,
		TotalFood:     domain.TotalFood,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
