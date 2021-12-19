package users

import (
	"Daily-Calorie-App-API/businesses/users"
	"Daily-Calorie-App-API/drivers/mysql/personaldata"
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

func toDomain(record Users) users.Domain {
	return users.Domain{
		ID:             int(record.ID),
		Name:           record.Name,
		Email:          record.Email,
		Password:       record.Password,
		AvatarUrl:      record.AvatarUrl,
		Gender:         record.Gender,
		PersonalDataID: int(record.PersonalDataID),
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		Model: gorm.Model{
			ID:        uint(domain.ID),
			CreatedAt: domain.CreatedAt,
			UpdatedAt: domain.UpdatedAt,
		},
		Name:           domain.Name,
		Email:          domain.Email,
		Password:       domain.Password,
		AvatarUrl:      domain.AvatarUrl,
		Gender:         domain.Gender,
		PersonalDataID: uint(domain.PersonalDataID),
	}
}

func toDomainArray(record []Users) []users.Domain {
	var result []users.Domain
	for _, v := range record {
		result = append(result, toDomain(v))
	}
	return result
}
