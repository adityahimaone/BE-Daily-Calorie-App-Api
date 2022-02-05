package personaldata

import (
	"Daily-Calorie-App-API/business/personaldata"
	"gorm.io/gorm"
)

type PersonalData struct {
	gorm.Model
	ID       uint    `gorm:"primary_key"`
	Calories float64 `gorm:"type:decimal(10,2)"`
	Weight   int     `gorm:"type:int"`
	Height   int     `gorm:"type:int"`
}

func toDomain(record PersonalData) personaldata.Domain {
	return personaldata.Domain{
		ID:       int(record.ID),
		Calories: record.Calories,
		Weight:   record.Weight,
		Height:   record.Height,
	}
}

func fromDomain(domain personaldata.Domain) PersonalData {
	return PersonalData{
		ID:       uint(domain.ID),
		Calories: domain.Calories,
		Weight:   domain.Weight,
		Height:   domain.Height,
	}
}
