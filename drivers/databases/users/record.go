package users

import (
	"Daily-Calorie-App-API/business/users"
	"Daily-Calorie-App-API/drivers/databases/personaldata"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID             uint   `gorm:"primary_key"`
	Name           string `gorm:"type:varchar(255);not null"`
	Email          string `gorm:"type:varchar(255);not null"`
	Password       string `gorm:"type:varchar(100);not null"`
	AvatarUrl      string `gorm:"type:varchar(255);not null"`
	Gender         string `gorm:"type:varchar(100);not null"`
	PersonalDataID uint
	PersonalData   personaldata.PersonalData `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (record *Users) toDomain() users.Domain {
	return users.Domain{
		ID:             int(record.ID),
		Name:           record.Name,
		Email:          record.Email,
		Password:       record.Password,
		AvatarUrl:      record.AvatarUrl,
		Gender:         record.Gender,
		PersonalDataID: int(record.PersonalDataID),
		Calories:       record.PersonalData.Calories,
		Weight:         record.PersonalData.Weight,
		Height:         record.PersonalData.Height,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		ID:             uint(domain.ID),
		Name:           domain.Name,
		Email:          domain.Email,
		Password:       domain.Password,
		AvatarUrl:      domain.AvatarUrl,
		Gender:         domain.Gender,
		PersonalDataID: uint(domain.PersonalDataID),
	}
}

func toDomainArray(record []Users) []users.Domain {
	var domains []users.Domain
	for _, v := range record {
		domains = append(domains, v.toDomain())
	}
	return domains
}
