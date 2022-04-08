package admins

import (
	"Daily-Calorie-App-API/business/admins"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	gorm.Model
	ID        int    `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(100);unique"`
	Password  string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (record *Admin) toDomain() admins.Domain {
	return admins.Domain{
		ID:        record.ID,
		Username:  record.Username,
		Password:  record.Password,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain admins.Domain) Admin {
	return Admin{
		ID:        domain.ID,
		Username:  domain.Username,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
