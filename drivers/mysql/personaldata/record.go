package personaldata

import (
	"Daily-Calorie-App-API/businesses/personaldata"
	"gorm.io/gorm"
)

type PersonalData struct {
	gorm.Model
	ID      uint    `gorm:"primary_key"`
	Calorie float64 `gorm:"type:decimal(10,2)"`
	Weight  int     `gorm:"type:int"`
	Height  int     `gorm:"type:int"`
}

func toDomain(record PersonalData) personaldata.Domain {
	return personaldata.Domain{
		ID:      int(record.ID),
		Calorie: record.Calorie,
		Weight:  record.Weight,
		Height:  record.Height,
	}
}

func fromDomain(domain personaldata.Domain) PersonalData {
	return PersonalData{
		ID:      uint(domain.ID),
		Calorie: domain.Calorie,
		Weight:  domain.Weight,
		Height:  domain.Height,
	}
}
